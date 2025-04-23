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

type CategoryDashboardController struct {
	CategoryUseCase usecase.CategoryUseCase
}

func NewCategoryController(categoryUC usecase.CategoryUseCase) *CategoryDashboardController {
	return &CategoryDashboardController{CategoryUseCase: categoryUC}
}

func (ctrl *CategoryDashboardController) CreateCategory(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	var reqCategory request.ReqCategory
	if err := c.BodyParser(&reqCategory); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqCategory, request.ReqCategoryErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	if err := ctrl.CategoryUseCase.CreateCategory(ctx, &reqCategory); err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed create category")
	}

	return response.SetResponseOK(c, "success create category", nil)
}

func (ctrl *CategoryDashboardController) GetCategoryByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	res, err := ctrl.CategoryUseCase.GetCategoryByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "Failed get category", err)
		return response.SetResponseBadRequest(c, "Failed get category", err)
	}

	return response.SetResponseOK(c, "success get category", res)
}

func (ctrl *CategoryDashboardController) GetListCategory(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	res, err := ctrl.CategoryUseCase.GetListCategory(ctx, utils.GetFiltersAndPagination(c))
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed get list category")
	}

	return response.SetResponseOK(c, "success get list user", res)
}

func (ctrl *CategoryDashboardController) UpdateCategoryByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)
	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqUpdate := request.ReqCategoryUpdate{}
	reqUpdate.ID = id
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqCategoryUpdateErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request ", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	res, err := ctrl.CategoryUseCase.UpdateCategoryByID(ctx, &reqUpdate)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed update category")
	}

	return response.SetResponseOK(c, "success update category", res)
}

func (ctrl *CategoryDashboardController) DeleteCategoryByID(c *fiber.Ctx) error {
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

	err = ctrl.CategoryUseCase.DeleteCategoryByID(ctx, id, reqData)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed delete category")
	}

	return response.SetResponseOK(c, "success delete category", nil)
}
