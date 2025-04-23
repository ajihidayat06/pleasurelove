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

type ProductDashboardController struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProductController(productUC usecase.ProductUseCase) *ProductDashboardController {
	return &ProductDashboardController{ProductUseCase: productUC}
}

func (ctrl *ProductDashboardController) CreateProduct(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	var reqProduct request.ReqProduct
	if err := c.BodyParser(&reqProduct); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqProduct, request.ReqProductErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	if err := ctrl.ProductUseCase.CreateProduct(ctx, &reqProduct); err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed create product")
	}

	return response.SetResponseOK(c, "success create product", nil)
}

func (ctrl *ProductDashboardController) GetProductByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	res, err := ctrl.ProductUseCase.GetProductByID(ctx, id)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed get product")
	}

	return response.SetResponseOK(c, "success get product", res)
}

func (ctrl *ProductDashboardController) GetListProduct(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	res, err := ctrl.ProductUseCase.GetListProduct(ctx, utils.GetFiltersAndPagination(c))
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed get list product")
	}

	return response.SetResponseOK(c, "success get list product", res)
}

func (ctrl *ProductDashboardController) UpdateProductByID(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)

	id, err := utils.ReadRequestParamID(c)
	if err != nil {
		logger.Error(ctx, "Failed get param id", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	reqUpdate := request.ReqProductUpdate{}
	reqUpdate.ID = id
	if err := c.BodyParser(&reqUpdate); err != nil {
		logger.Error(ctx, "Failed to parse request body", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	ok, errMsg := utils.ValidateRequest(reqUpdate, request.ReqProductUpdateErrorMessage)
	if !ok {
		err := fmt.Errorf("%s", errMsg)
		logger.Error(ctx, "error validate request", err)
		return response.SetResponseBadRequest(c, "Invalid request", err)
	}

	res, err := ctrl.ProductUseCase.UpdateProductByID(ctx, &reqUpdate)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed update product")
	}

	return response.SetResponseOK(c, "success update product", res)
}

func (ctrl *ProductDashboardController) DeleteProductByID(c *fiber.Ctx) error {
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

	err = ctrl.ProductUseCase.DeleteProductByID(ctx, id, reqData)
	if err != nil {
		return errorutils.HandleUsecaseError(c, err, "Failed delete product")
	}

	return response.SetResponseOK(c, "success delete product", nil)
}
