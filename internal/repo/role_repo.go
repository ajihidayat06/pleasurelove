package repo

import (
	"context"
	"pleasurelove/internal/models"
	"time"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(ctx context.Context, role *models.Roles) error
	GetRoleByID(ctx context.Context, id int64) (models.Roles, error)
	GetListRole(ctx context.Context, listStruct *models.GetListStruct) ([]models.Roles, int64, error)
	UpdateRoleByID(ctx context.Context, id int64, updatedAt time.Time, role models.Roles) (models.Roles, error)
	DeleteRoleByID(ctx context.Context, id int64, updatedAt time.Time) error
	GetRoleByCode(ctx context.Context, code string) (models.Roles, error)
}

type roleRepository struct {
	AbstractRepo
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		AbstractRepo: AbstractRepo{
			db: db,
		},
	}
}

func (r *roleRepository) Create(ctx context.Context, role *models.Roles) error {
	return r.getDB(ctx).WithContext(ctx).Create(role).Error
}

func (r *roleRepository) GetRoleByID(ctx context.Context, id int64) (models.Roles, error) {
	var role models.Roles

	err := r.db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Preload("RolePermissions").
		Preload("RolePermissions.Permissions").
		Where(" id = ? ", id).
		First(&role).Error
	if err != nil {
		return models.Roles{}, err
	}

	return role, nil
}

func (r *roleRepository) GetListRole(ctx context.Context, listStruct *models.GetListStruct) ([]models.Roles, int64, error) {
	var users []models.Roles
	var total int64

	// Query untuk menghitung total data tanpa Preload atau Pagination
	err := r.db.WithContext(ctx).
		Model(&models.Roles{}).
		Scopes(r.withCheckScope(ctx), r.applyFilters(listStruct.Filters)).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Query untuk mendapatkan data dengan Pagination dan Sorting
	err = r.db.WithContext(ctx).
		Model(&models.Roles{}).
		Scopes(r.withCheckScope(ctx), r.applyFiltersAndPaginationAndOrder(listStruct)).
		Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *roleRepository) UpdateRoleByID(ctx context.Context, id int64, updatedAt time.Time, role models.Roles) (models.Roles, error) {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Model(&role).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Save(&role).Error
	if err != nil {
		return models.Roles{}, err
	}
	return role, nil
}

func (r *roleRepository) DeleteRoleByID(ctx context.Context, id int64, updatedAt time.Time) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Delete(&models.Roles{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) GetRoleByCode(ctx context.Context, code string) (models.Roles, error) {
	var role models.Roles
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&role).Error
	if err != nil {
		return models.Roles{}, err
	}
	return role, nil
}
