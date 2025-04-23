package repo

import (
	"context"
	"pleasurelove/internal/models"
	"time"

	"gorm.io/gorm"
)

type RolePermissionsRepository interface {
	Create(ctx context.Context, rolePermission []models.RolePermissions) ([]models.RolePermissions, error)
	GetRolePermissionByID(ctx context.Context, id int64) (models.RolePermissions, error)
	GetListRolePermissions(ctx context.Context) ([]models.RolePermissions, error)
	UpdateRolePermissionsByID(ctx context.Context, id int64, updatedAt time.Time, rolePermission models.RolePermissions) (models.RolePermissions, error)
	DeleteRolePermissionsByID(ctx context.Context, id int64, updatedAt time.Time) error
	DeleteRolePermissionsByRoleID(ctx context.Context, roleID int64) error
	UpdateRolePermissionsBulk(ctx context.Context, rolePermission []models.RolePermissions) ([]models.RolePermissions, error)
}

type rolePermissionsRepository struct {
	AbstractRepo
}

func NewRolePermissionsRepository(db *gorm.DB) RolePermissionsRepository {
	return &rolePermissionsRepository{
		AbstractRepo: AbstractRepo{
			db: db,
		},
	}
}

func (r *rolePermissionsRepository) Create(ctx context.Context, rolePermissions []models.RolePermissions) ([]models.RolePermissions, error) {
	err := r.getDB(ctx).WithContext(ctx).Create(&rolePermissions).Error
	if err != nil {
		return nil, err
	}

	return rolePermissions, nil
}

func (r *rolePermissionsRepository) GetRolePermissionByID(ctx context.Context, id int64) (models.RolePermissions, error) {
	var rolePermission models.RolePermissions
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&rolePermission).Error
	if err != nil {
		return models.RolePermissions{}, err
	}
	return rolePermission, nil
}

func (r *rolePermissionsRepository) GetListRolePermissions(ctx context.Context) ([]models.RolePermissions, error) {
	var rolePermissions []models.RolePermissions
	err := r.db.WithContext(ctx).Find(&rolePermissions).Error
	if err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (r *rolePermissionsRepository) UpdateRolePermissionsByID(ctx context.Context, id int64, updatedAt time.Time, rolePermission models.RolePermissions) (models.RolePermissions, error) {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Model(&rolePermission).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Updates(rolePermission).Error
	if err != nil {
		return models.RolePermissions{}, err
	}
	return rolePermission, nil
}

func (r *rolePermissionsRepository) DeleteRolePermissionsByID(ctx context.Context, id int64, updatedAt time.Time) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Delete(&models.RolePermissions{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *rolePermissionsRepository) DeleteRolePermissionsByRoleID(ctx context.Context, roleID int64) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Where("role_id = ?", roleID).
		Delete(&models.RolePermissions{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *rolePermissionsRepository) UpdateRolePermissionsBulk(ctx context.Context, rolePermission []models.RolePermissions) ([]models.RolePermissions, error) {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Save(&rolePermission).Error
	if err != nil {
		return []models.RolePermissions{}, err
	}
	return rolePermission, nil
}
