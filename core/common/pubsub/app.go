package pubsub

import (
	"context"
	"gojumpstart/core/service"
	"os"
	"sync"

	cloud_pubsub "cloud.google.com/go/pubsub"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type PubSubApp struct {
	// cloud pubsub client
	client *cloud_pubsub.Client
	// project id
	projectID string
	// serviceAccountFile
	serviceAccountFile string
	// zap Logger
	Logger *zap.Logger
	// Context
	ctx context.Context
	// core service
	Svc *service.Service
	// routers
	routers []*PubSubRouter
}

func NewPubSubApp(service *service.Service) *PubSubApp {
	appEnv := os.Getenv("APP_ENV")
	Logger, _ := zap.NewDevelopment()
	if appEnv == "prod" {
		Logger, _ = zap.NewProduction()
	}
	return &PubSubApp{
		Svc:    service,
		Logger: Logger,
	}
}

func (p *PubSubApp) initEnv() {
	p.ctx = context.Background()
	p.projectID = os.Getenv("GCP_PROJECT_ID")
	p.serviceAccountFile = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	if p.projectID == "" {
		p.Logger.Panic("GCP_PROJECT_ID is required")
	}
}

func (p *PubSubApp) initClient() {
	var client *cloud_pubsub.Client
	var err error

	if p.client != nil {
		return
	}
	if p.serviceAccountFile == "" {
		client, err = cloud_pubsub.NewClient(p.ctx, p.projectID)
		if err != nil {
			p.Logger.Panic("Error creating pubsub client", zap.Error(err))
		}
	} else {
		client, err = cloud_pubsub.NewClient(p.ctx, p.projectID, option.WithCredentialsFile(p.serviceAccountFile))
		if err != nil {
			p.Logger.Panic("Error creating pubsub client", zap.Error(err))
		}
	}
	p.client = client

}

func (p *PubSubApp) subscribe() {
	var wg sync.WaitGroup
	// wg.Add(1)
	// go func(p *PubSubApp, wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	// Create a new subscription with the given name.
	// 	cctx, cancel := context.WithCancel(p.ctx)
	// 	defer cancel()
	// 	sub1 := p.client.Subscription("first-sub")
	// 	err := sub1.Receive(cctx, func(ctx context.Context, msg *cloud_pubsub.Message) {
	// 		isoDateTime := time.Now().Format(time.RFC3339)
	// 		fmt.Printf("(%v)[first-sub] Got message: %q\n", isoDateTime, string(msg.Data))
	// 		msg.Ack()
	// 	})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }(p, &wg)
	// wg.Add(1)
	// go func(p *PubSubApp, wg *sync.WaitGroup) {
	// 	wg.Add(1)
	// 	defer wg.Done()
	// 	// Create a new subscription with the given name.
	// 	cctx, cancel := context.WithCancel(p.ctx)
	// 	defer cancel()
	// 	sub2 := p.client.Subscription("second-sub")
	// 	err := sub2.Receive(cctx, func(ctx context.Context, msg *cloud_pubsub.Message) {
	// 		var isoDateTime string = time.Now().Format(time.RFC3339)
	// 		fmt.Printf("(%v)[second-sub] Got message: %q\n", isoDateTime, string(msg.Data))
	// 		msg.Ack()
	// 	})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }(p, &wg)

	for _, router := range p.routers {
		wg.Add(1)
		go func(router *PubSubRouter, wg *sync.WaitGroup) {
			defer wg.Done()
			router.Listen()
		}(router, &wg)
	}
	wg.Wait()
}

func (p *PubSubApp) AddRoute(topic *PubSubTopic, subs *PubSubSubcription, handler func(context.Context, *cloud_pubsub.Message) error) {
	router := NewRouter(p, topic, subs, handler)
	p.routers = append(p.routers, router)

}

func (p *PubSubApp) close() {
	if p.client != nil {
		p.client.Close()
	}
}

func (p *PubSubApp) Run() {
	p.initEnv()
	p.initClient()
	defer p.close()
	p.subscribe()
}
