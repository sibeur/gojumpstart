package http

import (
	"gojumpstart/apps/http/handler"
	"gojumpstart/core/service"
	"os"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

// FiberApp represents a Fiber application.
type FiberApp struct {
	Instance    *fiber.App
	Svc         *service.Service
	userHandler *handler.UserHandler
	todoHandler *handler.TodoHandler
}

// NewFiberApp creates a new instance of FiberApp.
func NewFiberApp(service *service.Service) *FiberApp {
	instance := fiber.New()
	return &FiberApp{
		Instance:    instance,
		Svc:         service,
		userHandler: handler.NewUserHandler(instance, service),
		todoHandler: handler.NewTodoHandler(instance, service),
	}
}

// beforeMiddlewares sets up the middlewares to be executed before the main request handler.
func (f *FiberApp) beforeMiddlewares() {
	appEnv := os.Getenv("APP_ENV")

	// Create a zap logger
	logger, _ := zap.NewDevelopment()
	if appEnv == "prod" {
		logger, _ = zap.NewProduction()
	}

	// Add fiberzap middleware
	f.Instance.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	// Add recover middleware
	f.Instance.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
}

// afterMiddlewares sets up the middlewares to be executed after the main request handler.
func (f *FiberApp) afterMiddlewares() {

}

// Run starts the Fiber application and listens for incoming requests.
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
