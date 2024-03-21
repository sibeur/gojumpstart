package handler

import (
	"gojumpstart/core/entity"
	"gojumpstart/core/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	fiberInstance *fiber.App
	svc           *service.Service
}

func NewUserHandler(fiberInstance *fiber.App, svc *service.Service) *UserHandler {
	return &UserHandler{
		fiberInstance: fiberInstance,
		svc:           svc,
	}
}

func (h *UserHandler) Router() {
	h.fiberInstance.Get("/users", h.FindAllUsers)
	h.fiberInstance.Get("/user-create", h.CreateUser)
}

func (h *UserHandler) FindAllUsers(c *fiber.Ctx) error {
	users, err := h.svc.User.FindAll()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := &entity.User{
		Username: "username",
		Email:    "email",
		Password: "password",
	}

	err := h.svc.User.Create(user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}
