package errors

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx) error {

	code := fiber.ErrInternalServerError
	message := "Internal Server Error"
}