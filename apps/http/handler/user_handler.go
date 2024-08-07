package handler

import (
	"gojumpstart/apps/http/handler/dto"
	"gojumpstart/core/common"
	"gojumpstart/core/common/helper"
	"gojumpstart/core/entity"
	"gojumpstart/core/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	fiberInstance *fiber.App
	svc           *service.Service
}

func NewUserHandler(fiberInstance *fiber.App, svc *service.Service) *UserHandler {
	return &UserHandler{
		fiberInstance: fiberInstance,
		svc:           svc,
	}
}

func (h *UserHandler) Router() {
	user := h.fiberInstance.Group("/users")
	user.Get("/", h.findAllUsers)
	user.Get("/create", h.createUser)
	user.Post("/", h.createUserWithValidator)
}

func (h *UserHandler) findAllUsers(c *fiber.Ctx) error {
	users, err := h.svc.User.FindAll()
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil, nil)
	}

	return common.SuccessResponse(c, "Success", users, nil)
}

func (h *UserHandler) createUser(c *fiber.Ctx) error {
	user := &entity.User{
		Username: "username",
		Email:    "email",
		Password: "password",
	}

	err := h.svc.User.Create(user)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil, nil)
	}

	return common.SuccessResponse(c, "", user, nil)
}

func (h *UserHandler) createUserWithValidator(c *fiber.Ctx) error {
	userData := new(dto.UserDTO)

	if err := c.BodyParser(userData); err != nil {
		return common.ErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil, nil)
	}

	validatorErrs := helper.NewErrValidationMessageBuilder()
	fValidator := common.NewFiberValidator()

	if errs := fValidator.Validate(userData); len(errs) > 0 {
		validatorErrs.AddBulk(errs)
	}

	if validatorErrs.HasError() {
		return validatorErrs.Build(c)
	}

	user := &entity.User{
		Username: userData.Username,
		Email:    userData.Email,
		Password: userData.Password,
	}
	err := h.svc.User.Create(user)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil, nil)
	}

	return common.SuccessResponse(c, "", user, nil)
}
