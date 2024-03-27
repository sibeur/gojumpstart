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
	todos := h.fiberInstance.Group("/todos")
	todos.Get("/", h.findAllTodos)
	todos.Get("/create", h.createTodo)
}

func (h *TodoHandler) findAllTodos(c *fiber.Ctx) error {
	todos, err := h.svc.Todo.FindAll()
	if err != nil {
		return errorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return successResponse(c, "", todos)
}

func (h *TodoHandler) createTodo(c *fiber.Ctx) error {
	todo := &entity.Todo{
		Title: "Foo",
		Desc:  "Bar",
		Done:  false,
	}

	err := h.svc.Todo.Create(todo)
	if err != nil {
		return errorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return successResponse(c, "", todo)
}
