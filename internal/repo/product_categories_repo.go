package repo

import (
	"context"
	"pleasurelove/internal/models"

	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	CreateBulk(ctx context.Context, pc []models.ProductCategory) error
	GetProductCategoryByProductID(ctx context.Context, productID int64) ([]models.ProductCategory, error)
	DeleteProductCategoryByProductID(ctx context.Context, productID int64) error
}

type productCategoryRepository struct {
	AbstractRepo
}

var (
	FilterProductCategory = map[string]string{
		"name": "name",
		"code": "code",
	}
	JoinsProductCategory                   = map[string]string{}
	ConstraintErrorMessagesProductcategory = map[string]string{
		"unique_product_categories_product_id_categories_id": "product dengan kategori ini sudah ada",
	}
)

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepository{
		AbstractRepo: AbstractRepo{
			db:              db,
			FilterAlias:     FilterProductCategory,
			Joins:           JoinsCategory,
			ConstraintError: ConstraintErrorMessagesProductcategory,
		},
	}
}

func (r *productCategoryRepository) CreateBulk(ctx context.Context, pc []models.ProductCategory) error {
	return r.getDB(ctx).WithContext(ctx).Create(pc).Error
}

func (r *productCategoryRepository) GetProductCategoryByProductID(ctx context.Context, productID int64) ([]models.ProductCategory, error) {
	var pcs []models.ProductCategory
	err := r.db.WithContext(ctx).
		Where("product_id = ?", productID).
		Find(&pcs).Error
	if err != nil {
		return nil, err
	}
	return pcs, nil
}

func (r *productCategoryRepository) DeleteProductCategoryByProductID(ctx context.Context, productID int64) error {
	return r.getDB(ctx).WithContext(ctx).
		Where("product_id = ?", productID).
		Delete(&models.ProductCategory{}).Error
}
