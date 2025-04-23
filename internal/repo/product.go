package repo

import (
	"context"
	"pleasurelove/internal/models"
	"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	GetProductByID(ctx context.Context, id int64) (models.Product, error)
	GetListProduct(ctx context.Context, listStruct *models.GetListStruct) ([]models.Product, int64, error)
	UpdateProductByID(ctx context.Context, id int64, updatedAt time.Time, product models.Product) (models.Product, error)
	DeleteProductByID(ctx context.Context, id int64, updatedAt time.Time) error
	GetProductByCode(ctx context.Context, code string) (models.Product, error)
}

type productRepository struct {
	AbstractRepo
}

var (
	FilterProduct = map[string]string{
		"name": "name",
		"code": "code",
	}
	JoinsProduct                   = map[string]string{}
	ProductConstraintErrorMessages = map[string]string{
		"unique_product_code": "Kode produk sudah digunakan",
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		AbstractRepo: AbstractRepo{
			db:              db,
			FilterAlias:     FilterProduct,
			Joins:           JoinsProduct,
			ConstraintError: ProductConstraintErrorMessages,
		},
	}
}

func (r *productRepository) Create(ctx context.Context, product *models.Product) error {
	return r.getDB(ctx).WithContext(ctx).Create(product).Error
}

func (r *productRepository) GetProductByID(ctx context.Context, id int64) (models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Preload("ProductCategory").
		Preload("ProductCategory.Category").
		Where("id = ?", id).
		First(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *productRepository) GetListProduct(ctx context.Context, listStruct *models.GetListStruct) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	err := r.db.WithContext(ctx).
		Model(&models.Product{}).
		Scopes(r.applyFilters(listStruct.Filters)).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).
		Model(&models.Product{}).
		Scopes(r.applyFiltersAndPaginationAndOrder(listStruct)).
		Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepository) UpdateProductByID(ctx context.Context, id int64, updatedAt time.Time, product models.Product) (models.Product, error) {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Model(&product).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Updates(product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *productRepository) DeleteProductByID(ctx context.Context, id int64, updatedAt time.Time) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Delete(&models.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) GetProductByCode(ctx context.Context, code string) (models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).
		Where("code = ?", code).
		First(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
