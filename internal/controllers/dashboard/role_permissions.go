package dashboard

import (
	"fmt"
	"pleasurelove/internal/dto/request"
	"pleasurelove/internal/dto/response"
	"pleasurelove/internal/usecase"
	"pleasurelove/internal/utils"
	"pleasurelove/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type RolePermissionsController struct {
	RolePermissionsUseCase usecase.RolePermissionsUsecase
}

func NewRolePermissionsController(rolePermissionsUC usecase.RolePermissionsUsecase) *RolePermissionsController {
	return &RolePermissionsController{RolePermissionsUseCase: rolePermissionsUC}
}

func (ctrl *RolePermissionsController) CreateRolePermission(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	var reqRolePermission request.ReqRolePermission
	if err := c.BodyParser(&reqRolePermission); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqRolePermission, request.ReqRolePermissionErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	// if err := ctrl.RolePermissionsUseCase.CreateRolePermission(ctx, reqRolePermission); err != nil {
	// 	logger.Error(ctx, "Failed to create role permission", err)
	// 	return response.SetResponseBadRequest(c, "Failed to create role permission", err)
	// }

	return response.SetResponseOK(c, "success create role permission", nil)
}

func (ctrl *RolePermissionsController) GetRolePermissionByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	res, err := ctrl.RolePermissionsUseCase.GetRolePermissionByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "Failed get role permission", err)
		return response.SetResponseBadRequest(c, "Failed get role permission", err)
	}

	return response.SetResponseOK(c, "success get role permission", res)
}

func (ctrl *RolePermissionsController) GetListRolePermissions(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	res, err := ctrl.RolePermissionsUseCase.GetListRolePermissions(ctx)
	if err != nil {
		logger.Error(ctx, "Failed get list role permissions", err)
		return response.SetResponseBadRequest(c, "Failed get list role permissions", err)
	}

	return response.SetResponseOK(c, "success get list role permissions", res)
}

func (ctrl *RolePermissionsController) UpdateRolePermissionByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqUpdate := request.ReqRolePermission{}
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqRolePermissionErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	res, err := ctrl.RolePermissionsUseCase.UpdateRolePermissionByID(ctx, id, reqUpdate.UpdatedAt, reqUpdate)
	if err != nil {
		logger.Error(ctx, "Failed update role permission", err)
		return response.SetResponseBadRequest(c, "Failed update role permission", err)
	}

	return response.SetResponseOK(c, "success update role permission", res)
}

func (ctrl *RolePermissionsController) DeleteRolePermissionByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqData := request.AbstractRequest{}
	if err := c.BodyParser(&reqData); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	err = ctrl.RolePermissionsUseCase.DeleteRolePermissionByID(ctx, id, reqData.UpdatedAt)
	if err != nil {
		logger.Error(ctx, "Failed delete role permission", err)
		return response.SetResponseBadRequest(c, "Failed delete role permission", err)
	}

	return response.SetResponseOK(c, "success delete role permission", nil)
}
