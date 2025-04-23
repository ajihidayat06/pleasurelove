package repo

import (
	"context"
	"pleasurelove/internal/models"
	"time"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *models.Category) error
	GetCategoryByID(ctx context.Context, id int64) (models.Category, error)
	GetListCategory(ctx context.Context, listStruct *models.GetListStruct) ([]models.Category, int64, error)
	UpdateCategoryByID(ctx context.Context, id int64, updatedAt time.Time, category models.Category) (models.Category, error)
	DeleteCategoryByID(ctx context.Context, id int64, updatedAt time.Time) error
	GetCategoryByNameOrCode(ctx context.Context, name string, code string) (models.Category, error)
	GetCategoryByListIDs(ctx context.Context, ids []int64) ([]models.Category, error)
}

type categoryRepository struct {
	AbstractRepo
}

var (
	FilterCategory = map[string]string{
		"name": "name",
	}
	JoinsCategory           = map[string]string{}
	ConstraintErrorMessages = map[string]string{
		"idx_categories_code": "Kode kategori sudah digunakan",
		"idx_categories_slug": "Nama kategori sudah digunakan",
	}
)

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		AbstractRepo: AbstractRepo{
			db:              db,
			FilterAlias:     FilterCategory,
			Joins:           JoinsCategory,
			ConstraintError: ConstraintErrorMessages,
		},
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *models.Category) error {
	return r.getDB(ctx).WithContext(ctx).Create(category).Error
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int64) (models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (r *categoryRepository) GetListCategory(ctx context.Context, listStruct *models.GetListStruct) ([]models.Category, int64, error) {
	var categories []models.Category
	var total int64

	err := r.db.WithContext(ctx).
		Model(&models.Category{}).
		Scopes(r.applyFilters(listStruct.Filters)).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).
		Model(&models.Category{}).
		Scopes(r.applyFiltersAndPaginationAndOrder(listStruct)).
		Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *categoryRepository) UpdateCategoryByID(ctx context.Context, id int64, updatedAt time.Time, category models.Category) (models.Category, error) {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Model(&category).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Updates(category).Error
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (r *categoryRepository) DeleteCategoryByID(ctx context.Context, id int64, updatedAt time.Time) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Delete(&models.Category{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) GetCategoryByNameOrCode(ctx context.Context, name string, code string) (models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).
		Where("name = ? OR code = ?", name, code).
		First(&category).Error
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (r *categoryRepository) GetCategoryByListIDs(ctx context.Context, ids []int64) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
