package pubsub

import (
	"context"

	cloud_pubsub "cloud.google.com/go/pubsub"
	"go.uber.org/zap"
)

type PubSubTopic struct {
	Name    string
	topic   *cloud_pubsub.Topic
	Options *cloud_pubsub.TopicConfig
}

type PubSubSubcription struct {
	Name         string
	subscription *cloud_pubsub.Subscription
	Options      *cloud_pubsub.SubscriptionConfig
}

type PubSubRouteHandler func(ctx context.Context, msg *cloud_pubsub.Message) error

type PubSubRouter struct {
	App          *PubSubApp
	Topic        *PubSubTopic
	Subscription *PubSubSubcription
	Handler      PubSubRouteHandler
}

func NewRouter(app *PubSubApp, topic *PubSubTopic, subscription *PubSubSubcription, handler PubSubRouteHandler) *PubSubRouter {
	return &PubSubRouter{
		App:          app,
		Topic:        topic,
		Subscription: subscription,
		Handler:      handler,
	}
}

func (r *PubSubRouter) CreateOrGetTopic() error {
	var err error
	topic := r.App.client.Topic(r.Topic.Name)
	exist, err := topic.Exists(r.App.ctx)
	if err != nil {
		return err
	}

	if exist {
		r.Topic.topic = topic
		return nil
	}

	options := r.Topic.Options
	if options == nil {
		topic, err = r.App.client.CreateTopic(r.App.ctx, r.Topic.Name)
		if err != nil {
			return err
		}
	} else {
		topic, err = r.App.client.CreateTopicWithConfig(r.App.ctx, r.Topic.Name, options)
		if err != nil {
			return err
		}
	}

	r.Topic.topic = topic
	return nil
}

func (r *PubSubRouter) CreateOrGetSubscription() error {
	options := r.Subscription.Options
	if options == nil {
		options = &cloud_pubsub.SubscriptionConfig{}
	}
	options.Topic = r.Topic.topic

	subscription := r.App.client.Subscription(r.Subscription.Name)
	exist, err := subscription.Exists(r.App.ctx)
	if err != nil {
		return err
	}

	if exist {
		r.Subscription.subscription = subscription
		return nil
	}

	subscription, err = r.App.client.CreateSubscription(r.App.ctx, r.Subscription.Name, *options)
	if err != nil {
		return err
	}
	r.Subscription.subscription = subscription
	return nil
}

func (r *PubSubRouter) Init() error {
	if err := r.CreateOrGetTopic(); err != nil {
		return err
	}
	if err := r.CreateOrGetSubscription(); err != nil {
		return err
	}
	return nil
}

func (r *PubSubRouter) Publish(msg *cloud_pubsub.Message) error {
	_, err := r.Topic.topic.Publish(r.App.ctx, msg).Get(r.App.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *PubSubRouter) Listen() {
	err := r.Init()
	if err != nil {
		r.App.Logger.Error("error initializing pubsub router", zap.Error(err))
		return
	}
	cctx, cancel := context.WithCancel(r.App.ctx)
	defer cancel()
	r.App.Logger.Info("listening for messages", zap.String("subscription", r.Subscription.subscription.String()))
	err = r.Subscription.subscription.Receive(cctx, func(ctx context.Context, msg *cloud_pubsub.Message) {
		if err := r.Handler(ctx, msg); err != nil {
			r.App.Logger.Error("error processing message", zap.Error(err))
			msg.Nack()
		}
		msg.Ack()
	})

	if err != nil {
		r.App.Logger.Error("error receiving message", zap.Error(err))
	}
}
