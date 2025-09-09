package models
import "time"
type Audit struct{
   User     string    `json:"user"`
    API      string    `json:"api"`
    DateTime time.Time `json:"dateTime"`
    IP       string    `json:"ip"`
    Message  string    `json:"message"`
    TraceID  string    `json:"traceId"`
//    Role     string    `json:"role"`
    Action   string    `json:"action"`
    Type     string    `json:"type"`
}