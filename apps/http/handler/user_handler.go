package handler

import (
	"gojumpstart/apps/http/handler/dto"
	"gojumpstart/core/common"
	"gojumpstart/core/common/helper"
	"gojumpstart/core/entity"
	"gojumpstart/core/service"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
	user.Post("/", h.createUser)
}

func (h *UserHandler) findAllUsers(c *fiber.Ctx) error {
	page := int64(c.QueryInt("page", 1))
	perPage := int64(c.QueryInt("per_page", 10))
	filter := &entity.UserListFilter{
		Search: c.Query("search", ""),
	}

	paginateData, err := h.svc.User.FindAllPaginate(page, perPage, filter)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil, nil)
	}

	response := make([]map[string]any, 0)

	for _, user := range paginateData.Data.([]*entity.User) {
		response = append(response, user.ToJSON())
	}
	return common.SuccessResponse(c, "Success", response, paginateData.Meta)
}

func (h *UserHandler) createUser(c *fiber.Ctx) error {
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

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil, nil)
	}

	user := &entity.User{
		Username: userData.Username,
		Email:    userData.Email,
		Password: string(bcryptPassword),
	}

	err = h.svc.User.Create(user)
	if err != nil {
		return common.ErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil, nil)
	}

	return common.SuccessResponse(c, "", user.ToJSON(), nil)
}
