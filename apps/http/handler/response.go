package handler

import (
	"gojumpstart/core/common"

	"github.com/gofiber/fiber/v2"
)

func successResponse(c *fiber.Ctx, message string, data any, meta any) error {
	jsonResp := fiber.Map{"message": message, "data": data, "meta": meta, "errors": nil}
	return c.JSON(jsonResp)
}

func errorResponse(c *fiber.Ctx, httpStatus int, message string, errorMessages []common.FiberErrorMessage, data any) error {
	jsonResp := fiber.Map{"message": message, "data": data, "meta": nil, "errors": errorMessages}
	return c.Status(httpStatus).JSON(jsonResp)
}
