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

type PermissionController struct {
	PermissionUseCase usecase.PermissionUseCase
}

func NewPermissionController(permissionUC usecase.PermissionUseCase) *PermissionController {
	return &PermissionController{PermissionUseCase: permissionUC}
}

func (ctrl *PermissionController) CreatePermission(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	var reqPermission request.ReqPermission
	if err := c.BodyParser(&reqPermission); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqPermission, request.ReqPermissionErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	if err := ctrl.PermissionUseCase.CreatePermission(ctx, &reqPermission); err != nil {
		logger.Error(ctx, "Failed to create permission", err)
		return response.SetResponseBadRequest(c, "Failed to create permission", err)
	}

	return response.SetResponseOK(c, "success create permission", nil)
}

func (ctrl *PermissionController) GetPermissionByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	res, err := ctrl.PermissionUseCase.GetPermissionByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "Failed get permission", err)
		return response.SetResponseBadRequest(c, "Failed get permission", err)
	}

	return response.SetResponseOK(c, "success get permission", res)
}

func (ctrl *PermissionController) GetListPermission(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	res, err := ctrl.PermissionUseCase.GetListPermission(ctx)
	if err != nil {
		logger.Error(ctx, "Failed get list permission", err)
		return response.SetResponseBadRequest(c, "Failed get list permission", err)
	}

	return response.SetResponseOK(c, "success get list permission", res)
}

func (ctrl *PermissionController) UpdatePermissionByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqUpdate := request.ReqPermissionUpdate{}
	reqUpdate.ID = id
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqPermissionUpdateErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	res, err := ctrl.PermissionUseCase.UpdatePermissionByID(ctx, &reqUpdate)
	if err != nil {
		logger.Error(ctx, "Failed update permission", err)
		return response.SetResponseBadRequest(c, "Failed update permission", err)
	}

	return response.SetResponseOK(c, "success update permission", res)
}

func (ctrl *PermissionController) DeletePermissionByID(c *fiber.Ctx) error {
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

	err = ctrl.PermissionUseCase.DeletePermissionByID(ctx, id, reqData.UpdatedAt)
	if err != nil {
		logger.Error(ctx, "Failed delete permission", err)
		return response.SetResponseBadRequest(c, "Failed delete permission", err)
	}

	return response.SetResponseOK(c, "success delete permission", nil)
}
