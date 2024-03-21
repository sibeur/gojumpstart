package pubsub

import "gojumpstart/core/service"

type PubSubApp struct {
	// cloud pubsub client
	// Client *pubsub.Client
	// core service
	Svc *service.Service
	// handler
}

func NewPubSubApp(service *service.Service) *PubSubApp {
	return &PubSubApp{
		Svc: service,
	}
}

func (p *PubSubApp) Run() {
	// p.Client = pubsub.NewClient(context.Background(), "project-id")
	// defer p.Client.Close()
	// p.subscribe()
}
