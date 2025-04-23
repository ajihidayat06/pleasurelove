package usecase

import (
	"context"
	"errors"
	"pleasurelove/internal/constanta"
	"pleasurelove/internal/dto/request"
	"pleasurelove/internal/dto/response"
	"pleasurelove/internal/models"
	"pleasurelove/internal/repo"
	"pleasurelove/internal/utils"
	"pleasurelove/internal/utils/errorutils"
	"pleasurelove/pkg/logger"
	"time"

	"gorm.io/gorm"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, req *request.ReqProduct) error
	GetProductByID(ctx context.Context, id int64) (response.DetailProductResponse, error)
	GetListProduct(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.ProductResponse], error)
	UpdateProductByID(ctx context.Context, req *request.ReqProductUpdate) (response.ProductResponse, error)
	DeleteProductByID(ctx context.Context, id int64, reqData request.AbstractRequest) error
}

type productUseCase struct {
	db                  *gorm.DB
	productRepo         repo.ProductRepository
	categoryRepo        repo.CategoryRepository
	productCategoryRepo repo.ProductCategoryRepository
}

func NewProductUseCase(db *gorm.DB,
	productRepo repo.ProductRepository,
	categoryRepo repo.CategoryRepository,
	productCategoryRepo repo.ProductCategoryRepository) ProductUseCase {
	return &productUseCase{
		db:                  db,
		productRepo:         productRepo,
		categoryRepo:        categoryRepo,
		productCategoryRepo: productCategoryRepo,
	}
}

func (uc *productUseCase) CreateProduct(ctx context.Context, req *request.ReqProduct) error {
	err := req.ValidateRequestCreate()
	if err != nil {
		logger.Error(ctx, "Failed to validate request", err)
		return err
	}

	err = uc.validateCategory(ctx, req.CategoryID)
	if err != nil {
		return err
	}

	userID, err := utils.GetUserIDFromCtx(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to get user id from context", err)
		return errorutils.ErrDataNotFound
	}

	product := models.Product{
		Name:        req.Name,
		Code:        req.Code,
		Barcode:     req.Barcode,
		Description: req.Description,
		Brand:       req.Brand,
		Unit:        req.Unit,
		Price:       req.Price,
		CostPrice:   req.CostPrice,
		Discount:    req.Discount,
		IsActive:    req.IsActive,
		HasVarian:   req.HasVarian,
		CreatedBy:   userID,
		UpdatedBy:   userID,
	}

	//TODO: validate and add product_varian

	return processWithTx(ctx, uc.db, func(ctx context.Context) error {
		err := uc.productRepo.Create(ctx, &product)
		if err != nil {
			return errorutils.HandleRepoErrorWrite(ctx, err, repo.GetContraintErrMessage(uc.productRepo))
		}

		// insert product_categories
		if len(req.CategoryID) > 0 {
			var pcs []models.ProductCategory
			for _, v := range req.CategoryID {
				pcs = append(pcs, models.ProductCategory{
					ProductID:    product.ID,
					CategoriesID: v,
					CreatedBy:    userID,
					UpdatedBy:    userID,
				})
			}

			err = uc.productCategoryRepo.CreateBulk(ctx, pcs)
			if err != nil {
				return errorutils.HandleRepoErrorWrite(ctx, err, repo.GetContraintErrMessage(uc.productCategoryRepo))
			}
		}

		return nil
	})
}

func (uc *productUseCase) GetProductByID(ctx context.Context, id int64) (response.DetailProductResponse, error) {
	product, err := uc.productRepo.GetProductByID(ctx, id)
	if err != nil {
		return response.DetailProductResponse{}, errorutils.HandleRepoError(ctx, err)
	}
	return response.SetDetailProductResponse(product), nil
}

func (uc *productUseCase) GetListProduct(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.ProductResponse], error) {
	products, count, err := uc.productRepo.GetListProduct(ctx, listStruct)
	if err != nil {
		return response.ListResponse[response.ProductResponse]{}, errorutils.HandleRepoError(ctx, err)
	}

	listResponse := response.MapToListResponse(response.SetResponseListProduct(products), count, listStruct, repo.GetFilterAvailableFromRepo(uc.productRepo))
	return listResponse, nil
}

func (uc *productUseCase) UpdateProductByID(ctx context.Context, req *request.ReqProductUpdate) (response.ProductResponse, error) {
	if err := req.ValidateUpdatedAt(); err != nil {
		return response.ProductResponse{}, err
	}

	if err := req.ValidateRequestCreate(); err != nil {
		return response.ProductResponse{}, err
	}

	productDb, err := uc.productRepo.GetProductByID(ctx, req.ID)
	if err != nil {
		return response.ProductResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	if !utils.ValidateUpdatedAtRequest(req.UpdatedAt, productDb.UpdatedAt) {
		return response.ProductResponse{}, errorutils.ErrDataDataUpdated
	}

	validateUnique, err := uc.productRepo.GetProductByCode(ctx, req.Code)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.ProductResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	if validateUnique.ID != 0 && validateUnique.ID != req.ID {
		return response.ProductResponse{}, errorutils.HandleCustomError(ctx, err, errorutils.ErrMessaageDataAlreadyExists, constanta.FieldCode, constanta.FieldName)
	}

	isUpdateCategory := uc.validateUpdateCategoryData(ctx, req.CategoryID, *productDb.ProductCategory)

	if isUpdateCategory {
		err = uc.validateCategory(ctx, req.CategoryID)
		if err != nil {
			return response.ProductResponse{}, err
		}
	}

	userID, err := utils.GetUserIDFromCtx(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to get user id from context", err)
		return response.ProductResponse{}, errorutils.ErrDataNotFound
	}

	product := models.Product{
		ID:          req.ID,
		Name:        req.Name,
		Code:        req.Code,
		Barcode:     req.Barcode,
		Description: req.Description,
		Brand:       req.Brand,
		Unit:        req.Unit,
		Price:       req.Price,
		CostPrice:   req.CostPrice,
		Discount:    req.Discount,
		IsActive:    req.IsActive,
		HasVarian:   req.HasVarian,
		CreatedAt:   productDb.CreatedAt,
		CreatedBy:   productDb.CreatedBy,
		UpdatedAt:   time.Now(),
		UpdatedBy:   userID,
	}

	var updated models.Product
	err = processWithTx(ctx, uc.db, func(ctx context.Context) error {
		updated, err = uc.productRepo.UpdateProductByID(ctx, req.ID, req.UpdatedAt, product)
		if err != nil {
			return errorutils.HandleRepoErrorWrite(ctx, err, repo.GetContraintErrMessage(uc.productRepo))
		}

		if isUpdateCategory {
			err = uc.updateDataCategory(ctx, &product, req, userID)
			if err != nil {
				return errorutils.HandleRepoErrorWrite(ctx, err, repo.GetContraintErrMessage(uc.productRepo))
			}
		}

		return nil
	})

	if err != nil {
		return response.ProductResponse{}, err
	}

	return response.SetProductResponse(updated), nil
}

func (uc *productUseCase) DeleteProductByID(ctx context.Context, id int64, reqData request.AbstractRequest) error {
	if err := reqData.ValidateUpdatedAt(); err != nil {
		return err
	}

	product, err := uc.productRepo.GetProductByID(ctx, id)
	if err != nil {
		return errorutils.HandleRepoError(ctx, err)
	}

	if !utils.ValidateUpdatedAtRequest(reqData.UpdatedAt, product.UpdatedAt) {
		return errorutils.ErrDataDataUpdated
	}

	return processWithTx(ctx, uc.db, func(ctx context.Context) error {
		err = uc.productCategoryRepo.DeleteProductCategoryByProductID(ctx, product.ID)
		if err != nil {
			return errorutils.HandleRepoError(ctx, err)
		}

		err := uc.productRepo.DeleteProductByID(ctx, id, reqData.UpdatedAt)
		if err != nil {
			return errorutils.HandleRepoError(ctx, err)
		}
		return nil
	})
}

func (uc *productUseCase) validateCategory(ctx context.Context, req []int64) (err error) {
	var (
		categories []models.Category
	)
	if len(req) > 0 {
		categories, err = uc.categoryRepo.GetCategoryByListIDs(ctx, req)
		if err != nil {
			return errorutils.HandleRepoError(ctx, err)
		}
	}

	if len(req) != len(categories) {
		return errorutils.HandleCustomError(ctx, nil, errorutils.ErrMessageDataNotFound, constanta.FieldCategory)
	}

	return nil
}

func (uc *productUseCase) validateUpdateCategoryData(ctx context.Context, req []int64, categoryDb []models.ProductCategory) (isUpdateCategory bool) {
	if len(req) > 0 {
		count := 0
		for _, v := range req {
			for _, cat := range categoryDb {
				if v == cat.CategoriesID {
					break
				}
			}

			count++
		}

		if count != len(categoryDb) {
			return true
		}
	}

	return
}

func (uc *productUseCase) updateDataCategory(ctx context.Context, product *models.Product, req *request.ReqProductUpdate, userID int64) (err error) {
	err = uc.productCategoryRepo.DeleteProductCategoryByProductID(ctx, product.ID)
	if err != nil {
		return errorutils.HandleRepoErrorWrite(ctx, err, repo.GetContraintErrMessage(uc.productRepo))
	}

	var (
		pcs []models.ProductCategory
	)
	for _, v := range req.CategoryID {
		pcs = append(pcs, models.ProductCategory{
			ProductID:    product.ID,
			CategoriesID: v,
			CreatedBy:    userID,
			UpdatedBy:    userID,
		})
	}

	err = uc.productCategoryRepo.CreateBulk(ctx, pcs)
	if err != nil {
		return errorutils.HandleRepoErrorWrite(ctx, err, repo.GetContraintErrMessage(uc.productCategoryRepo))
	}

	return
}
