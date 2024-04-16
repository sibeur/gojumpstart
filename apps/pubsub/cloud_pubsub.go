package pubsub

import (
	"gojumpstart/apps/pubsub/handler"
	core_pubsub "gojumpstart/core/common/pubsub"
	"gojumpstart/core/service"
)

type App struct {
	Instance    *core_pubsub.PubSubApp
	Svc         *service.Service
	userHandler *handler.UserHandler
}

func NewApp(service *service.Service) *App {
	instance := core_pubsub.NewPubSubApp(service)
	return &App{
		Instance:    instance,
		Svc:         service,
		userHandler: handler.NewUserHandler(instance, service),
	}
}

func (a *App) Run() {
	a.userHandler.Router()
	a.Instance.Run()
}
