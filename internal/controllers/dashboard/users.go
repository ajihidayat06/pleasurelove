package dashboard

import (
	"fmt"
	"pleasurelove/internal/dto/request"
	"pleasurelove/internal/dto/response"
	"pleasurelove/internal/usecase"
	"pleasurelove/internal/utils"
	"pleasurelove/internal/utils/errorutils"
	"pleasurelove/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type UserDahboardController struct {
	UserDashboardUsecase usecase.UserUseCase
}

func NewUserDashboardController(
	userUC usecase.UserUseCase,
) *UserDahboardController {
	return &UserDahboardController{UserDashboardUsecase: userUC}
}

func (ctrl *UserDahboardController) CreateUserDashboard(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	var reqUser request.ReqUser
	if err := c.BodyParser(&reqUser); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	ok, errMsg := utils.ValidateRequest(reqUser, request.ReqUserErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	if err := ctrl.UserDashboardUsecase.CreateUserDashboard(ctx, &reqUser); err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed create user")
	}

	return response.SetResponseOK(c, "success register user", nil)
}

func (ctrl *UserDahboardController) GetUserByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	res, err := ctrl.UserDashboardUsecase.GetUserByID(ctx, id)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed get user")
	}

	return response.SetResponseOK(c, "success get user", res)
}

func (ctrl *UserDahboardController) GetListUser(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	res, err := ctrl.UserDashboardUsecase.GetListUser(ctx, utils.GetFiltersAndPagination(c))
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed get list user")
	}

	return response.SetResponseOK(c, "success get list user", res)
}

func (ctrl *UserDahboardController) UpdateUserByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	reqUpdate := request.ReqUserUpdate{}
	reqUpdate.ID = id
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqUserUpdateErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	res, err := ctrl.UserDashboardUsecase.UpdateUserByID(ctx, &reqUpdate)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed update user")
	}

	return response.SetResponseOK(c, "success update user", res)
}

func (ctrl *UserDahboardController) DeleteUserByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	reqData := request.AbstractRequest{}
	if err := c.BodyParser(&reqData); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	err = ctrl.UserDashboardUsecase.DeleteUserByID(ctx, id, reqData)
	if err != nil {
		logger.Error(ctx, "Failed delete user", err)
		return response.SetResponseBadRequest(c, "Failed delete user", err)
	}

	return response.SetResponseOK(c, "success delete user", nil)
}
