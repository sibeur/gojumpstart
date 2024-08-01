package common

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, message string, data any, meta any) error {
	jsonResp := fiber.Map{"message": message, "data": data, "meta": meta, "errors": nil}
	return c.JSON(jsonResp)
}

func ErrorResponse(c *fiber.Ctx, httpStatus int, message string, errorMessages []FiberErrorMessage, data any) error {
	jsonResp := fiber.Map{"message": message, "data": data, "meta": nil, "errors": errorMessages}
	return c.Status(httpStatus).JSON(jsonResp)
}
