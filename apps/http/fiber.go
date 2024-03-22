package http

import (
	"gojumpstart/apps/http/handler"
	"gojumpstart/core/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberApp struct {
	Instance    *fiber.App
	Svc         *service.Service
	userHandler *handler.UserHandler
	todoHandler *handler.TodoHandler
}

func NewFiberApp(service *service.Service) *FiberApp {
	instance := fiber.New()
	return &FiberApp{
		Instance:    instance,
		Svc:         service,
		userHandler: handler.NewUserHandler(instance, service),
		todoHandler: handler.NewTodoHandler(instance, service),
	}
}

func (f *FiberApp) beforeMiddlewares() {
	f.Instance.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
}

func (f *FiberApp) afterMiddlewares() {

}

func (f *FiberApp) Run() {
	f.beforeMiddlewares()
	f.Instance.Get("/", func(c *fiber.Ctx) error {
		panic("Error")
		return c.SendString("Hello, World!")
	})
	f.userHandler.Router()
	f.todoHandler.Router()
	f.afterMiddlewares()
	f.Instance.Listen(":3000")
}
