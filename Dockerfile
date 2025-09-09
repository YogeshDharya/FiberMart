# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git for ldflags versioning
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X main.version=$(git describe --tags --always || echo dev)" \
    -o fiber-mart .

# Install git + tzdata for timezone info
RUN apk add --no-cache git tzdata

# Final stage
FROM gcr.io/distroless/base

WORKDIR /app

# Copy required files
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/fiber-mart /app/fiber-mart

# Expose port
EXPOSE 8080

# Healthcheck via HTTP endpoint (if you add /healthz route in Fiber)
HEALTHCHECK --interval=30s --timeout=3s \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/healthz || exit 1

# Run app
CMD ["/app/fiber-mart"]
