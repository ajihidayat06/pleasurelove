package repo

import (
	"context"
	"pleasurelove/internal/models"
	"time"

	"gorm.io/gorm"
)

type PermissionsRepository interface {
	Create(ctx context.Context, permission *models.Permissions) error
	GetPermissionsByID(ctx context.Context, id int64) (models.Permissions, error)
	GetListPermissions(ctx context.Context) ([]models.Permissions, error)
	UpdatePermissionsByID(ctx context.Context, id int64, updatedAt time.Time, permissions models.Permissions) (models.Permissions, error)
	DeletePermissionsByID(ctx context.Context, id int64, updatedAt time.Time) error
	GetPermissionsByListID(ctx context.Context, ids []int64) ([]models.Permissions, error)
}

type permissionsRepository struct {
	AbstractRepo
}

func NewPermissionsRepository(db *gorm.DB) PermissionsRepository {
	return &permissionsRepository{
		AbstractRepo: AbstractRepo{
			db: db,
		},
	}
}

func (r *permissionsRepository) Create(ctx context.Context, permissions *models.Permissions) error {
	return r.getDB(ctx).WithContext(ctx).Create(permissions).Error
}

func (r *permissionsRepository) GetPermissionsByID(ctx context.Context, id int64) (models.Permissions, error) {
	var permissions models.Permissions
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&permissions).Error
	if err != nil {
		return models.Permissions{}, err
	}
	return permissions, nil
}

func (r *permissionsRepository) GetListPermissions(ctx context.Context) ([]models.Permissions, error) {
	var permissionss []models.Permissions
	err := r.db.WithContext(ctx).Find(&permissionss).Error
	if err != nil {
		return nil, err
	}
	return permissionss, nil
}

func (r *permissionsRepository) UpdatePermissionsByID(ctx context.Context, id int64, updatedAt time.Time, permissions models.Permissions) (models.Permissions, error) {
	db := r.getDB(ctx) // Gunakan DB dari context jika ada transaksi

	err := db.WithContext(ctx).
		Model(&permissions).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Updates(permissions).Error
	if err != nil {
		return models.Permissions{}, err
	}
	return permissions, nil
}

func (r *permissionsRepository) DeletePermissionsByID(ctx context.Context, id int64, updatedAt time.Time) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Delete(&models.Permissions{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *permissionsRepository) GetPermissionsByListID(ctx context.Context, ids []int64) ([]models.Permissions, error) {
	var permissions []models.Permissions

	err := r.db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Where("id IN ?", ids).
		Find(&permissions).Error

	if err != nil {
		return nil, err
	}

	return permissions, nil
}
