package handler

import "github.com/gofiber/fiber/v2"

func successResponse(c *fiber.Ctx, message string, data any) error {
	jsonResp := fiber.Map{"message": message, "data": data}
	return c.JSON(jsonResp)
}

func errorResponse(c *fiber.Ctx, httpStatus int, message string, data any) error {
	jsonResp := fiber.Map{"message": message, "data": data}
	return c.Status(httpStatus).JSON(jsonResp)
}
