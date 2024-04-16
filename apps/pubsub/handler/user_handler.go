package handler

import (
	"context"
	"encoding/json"
	"fmt"
	core_pubsub "gojumpstart/core/common/pubsub"
	"gojumpstart/core/entity"
	"gojumpstart/core/service"

	cloud_pubsub "cloud.google.com/go/pubsub"
)

type UserHandler struct {
	pubsubInstance *core_pubsub.PubSubApp
	svc            *service.Service
}

func NewUserHandler(pubsubInstance *core_pubsub.PubSubApp, svc *service.Service) *UserHandler {
	return &UserHandler{
		pubsubInstance: pubsubInstance,
		svc:            svc,
	}
}

func (h *UserHandler) Router() {
	h.pubsubInstance.AddRoute(&core_pubsub.PubSubTopic{
		Name: "create-user",
	}, &core_pubsub.PubSubSubcription{
		Name: "create-user-sub",
	}, h.createUser)
}

func (h *UserHandler) createUser(ctx context.Context, msg *cloud_pubsub.Message) error {
	var user entity.User
	data := msg.Data

	h.pubsubInstance.Logger.Info(fmt.Sprintf("Received message: %s", data))
	if err := json.Unmarshal(data, &user); err != nil {
		return err
	}
	// user := &entity.User{
	// 	Username: "username",
	// 	Email:    "email",
	// 	Password: "password",
	// }

	if err := h.svc.User.Create(&user); err != nil {
		return err
	}

	h.pubsubInstance.Logger.Info(fmt.Sprintf("User created: %v", user))
	return nil
}
