package http

import (
	"gojumpstart/apps/http/handler"
	"gojumpstart/core/service"

	"github.com/gofiber/fiber/v2"
)

type FiberApp struct {
	Instance    *fiber.App
	Svc         *service.Service
	userHandler *handler.UserHandler
}

func NewFiberApp(service *service.Service) *FiberApp {
	instance := fiber.New()
	return &FiberApp{
		Instance:    instance,
		Svc:         service,
		userHandler: handler.NewUserHandler(instance, service),
	}
}

func (f *FiberApp) middlewares() {
	// f.Instance.Use(func(c *fiber.Ctx) error {
	// 	c.Set("X-Request-ID", "123456")
	// 	return c.Next()
	// })
}

func (f *FiberApp) Run() {
	f.middlewares()
	f.userHandler.Router()
	f.Instance.Listen(":3000")
}
