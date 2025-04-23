package usecase

import (
	"context"
	"errors"
	"time"

	"pleasurelove/internal/dto/request"
	"pleasurelove/internal/models"
	"pleasurelove/internal/repo"
)

type RolePermissionsUsecase interface {
	// CreateRolePermission(ctx context.Context, req request.ReqRolePermission) error
	GetRolePermissionByID(ctx context.Context, id int64) (models.RolePermissions, error)
	GetListRolePermissions(ctx context.Context) ([]models.RolePermissions, error)
	UpdateRolePermissionByID(ctx context.Context, id int64, updatedAt time.Time, req request.ReqRolePermission) (models.RolePermissions, error)
	DeleteRolePermissionByID(ctx context.Context, id int64, updatedAt time.Time) error
}

type rolePermissionsUsecase struct {
	repo repo.RolePermissionsRepository
}

func NewRolePermissionsUsecase(repo repo.RolePermissionsRepository) RolePermissionsUsecase {
	return &rolePermissionsUsecase{repo: repo}
}

// CreateRolePermission creates a new role-permission record.
// func (u *rolePermissionsUsecase) CreateRolePermission(ctx context.Context, req request.ReqRolePermission) error {
// 	// if req.RoleID <= 0 || req.PermissionID <= 0 {
// 	// 	return errors.New("invalid role ID or permission ID")
// 	// }

// 	rolePermission := models.RolePermissions{
// 		// RoleID:        req.RoleID,
// 		PermissionsID: req.PermissionID,
// 		AccessScope:   req.Scope,
// 		CreatedAt:     time.Now(),
// 		UpdatedAt:     time.Now(),
// 	}

// 	return u.repo.Create(ctx, &rolePermission)
// }

// GetRolePermissionByID retrieves a role-permission record by its ID.
func (u *rolePermissionsUsecase) GetRolePermissionByID(ctx context.Context, id int64) (models.RolePermissions, error) {
	if id <= 0 {
		return models.RolePermissions{}, errors.New("invalid ID")
	}

	return u.repo.GetRolePermissionByID(ctx, id)
}

// GetListRolePermissions retrieves all role-permission records.
func (u *rolePermissionsUsecase) GetListRolePermissions(ctx context.Context) ([]models.RolePermissions, error) {
	return u.repo.GetListRolePermissions(ctx)
}

// UpdateRolePermissionByID updates a role-permission record by its ID.
func (u *rolePermissionsUsecase) UpdateRolePermissionByID(ctx context.Context, id int64, updatedAt time.Time, req request.ReqRolePermission) (models.RolePermissions, error) {
	if id <= 0 {
		return models.RolePermissions{}, errors.New("invalid ID")
	}

	rolePermission := models.RolePermissions{
		// RoleID:        req.RoleID,
		PermissionsID: req.PermissionID,
		AccessScope:   req.Scope,
		UpdatedAt:     time.Now(),
	}

	return u.repo.UpdateRolePermissionsByID(ctx, id, updatedAt, rolePermission)
}

// DeleteRolePermissionByID deletes a role-permission record by its ID.
func (u *rolePermissionsUsecase) DeleteRolePermissionByID(ctx context.Context, id int64, updatedAt time.Time) error {
	if id <= 0 {
		return errors.New("invalid ID")
	}

	return u.repo.DeleteRolePermissionsByID(ctx, id, updatedAt)
}
