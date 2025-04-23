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

type RoleController struct {
	RoleUseCase usecase.RoleUseCase
}

func NewRoleController(roleUC usecase.RoleUseCase) *RoleController {
	return &RoleController{RoleUseCase: roleUC}
}

func (ctrl *RoleController) CreateRole(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	var reqRole request.ReqRoles
	if err := c.BodyParser(&reqRole); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	ok, errMsg := utils.ValidateRequest(reqRole, request.ReqRolesErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	if err := ctrl.RoleUseCase.CreateRole(ctx, &reqRole); err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed create role")
	}

	return response.SetResponseOK(c, "success create role", nil)
}

func (ctrl *RoleController) GetRoleByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	res, err := ctrl.RoleUseCase.GetRoleByID(ctx, id)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed get role")
	}

	return response.SetResponseOK(c, "success get role", res)
}

func (ctrl *RoleController) GetListRole(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	res, err := ctrl.RoleUseCase.GetListRole(ctx, utils.GetFiltersAndPagination(c))
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed get list role")
	}

	return response.SetResponseOK(c, "success get list role", res)
}

func (ctrl *RoleController) UpdateRoleByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	reqUpdate := request.ReqRoleUpdate{}
	reqUpdate.ID = id
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqRoleUpdateErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return response.SetResponseBadRequest(c, errorutils.ErrMessageInvalidRequestData, err)
	}

	res, err := ctrl.RoleUseCase.UpdateRoleByID(ctx, &reqUpdate)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed update role")
	}

	return response.SetResponseOK(c, "success update role", res)
}

func (ctrl *RoleController) DeleteRoleByID(c *fiber.Ctx) error {
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

	err = ctrl.RoleUseCase.DeleteRoleByID(ctx, id, reqData)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed delete role")
	}

	return response.SetResponseOK(c, "success delete role", nil)
}
