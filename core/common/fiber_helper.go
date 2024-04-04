package common

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type FiberErrorMessage struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Default error handler
var FiberDefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Return status code with error message
	return c.Status(code).JSON(fiber.Map{"message": err.Error(), "meta": nil, "data": nil, "errors": nil})
}
