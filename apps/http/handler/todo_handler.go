package handler

import (
	"gojumpstart/core/entity"
	"gojumpstart/core/service"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	fiberInstance *fiber.App
	svc           *service.Service
}

func NewTodoHandler(fiberInstance *fiber.App, svc *service.Service) *TodoHandler {
	return &TodoHandler{
		fiberInstance: fiberInstance,
		svc:           svc,
	}
}

func (h *TodoHandler) Router() {
	h.fiberInstance.Get("/todos", h.findAllTodos)
	h.fiberInstance.Get("/todo-create", h.createTodo)
}

func (h *TodoHandler) findAllTodos(c *fiber.Ctx) error {
	users, err := h.svc.Todo.FindAll()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(users)
}

func (h *TodoHandler) createTodo(c *fiber.Ctx) error {
	user := &entity.Todo{
		Title: "Foo",
		Desc:  "Bar",
		Done:  false,
	}

	err := h.svc.Todo.Create(user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}
