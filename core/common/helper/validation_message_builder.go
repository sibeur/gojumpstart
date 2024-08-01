package helper

import (
	"gojumpstart/core/common"

	"github.com/gofiber/fiber/v2"
)

type ErrValidationMessageBuilder struct {
	messages []common.FiberErrorMessage
}

func NewErrValidationMessageBuilder() *ErrValidationMessageBuilder {
	return &ErrValidationMessageBuilder{
		messages: []common.FiberErrorMessage{},
	}
}

func (b *ErrValidationMessageBuilder) Add(key string, value string) {
	b.messages = append(b.messages, common.FiberErrorMessage{
		Key:   key,
		Value: value,
	})
}

func (b *ErrValidationMessageBuilder) AddBulk(messages []common.FiberErrorMessage) {
	b.messages = append(b.messages, messages...)
}

func (b *ErrValidationMessageBuilder) HasError() bool {
	return len(b.messages) > 0
}

func (b *ErrValidationMessageBuilder) GetMessages() []common.FiberErrorMessage {
	return b.messages
}

func (b *ErrValidationMessageBuilder) Build(c *fiber.Ctx) error {
	return common.ErrorResponse(c, fiber.StatusBadRequest, "Validation error", b.messages, nil)
}
