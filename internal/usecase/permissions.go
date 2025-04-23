package usecase

import (
	"context"
	"pleasurelove/internal/dto/request"
	"pleasurelove/internal/models"
	"pleasurelove/internal/repo"
	"time"

	"gorm.io/gorm"
)

type PermissionUseCase interface {
	CreatePermission(ctx context.Context, req *request.ReqPermission) error
	GetPermissionByID(ctx context.Context, id int64) (models.Permissions, error)
	GetListPermission(ctx context.Context) ([]models.Permissions, error)
	UpdatePermissionByID(ctx context.Context, req *request.ReqPermissionUpdate) (models.Permissions, error)
	DeletePermissionByID(ctx context.Context, id int64, updatedAt time.Time) error
}

type permissionUseCase struct {
	db             *gorm.DB
	permissionRepo repo.PermissionsRepository
}

func NewPermissionUseCase(db *gorm.DB, permissionRepo repo.PermissionsRepository) PermissionUseCase {
	return &permissionUseCase{
		db:             db,
		permissionRepo: permissionRepo,
	}
}

func (uc *permissionUseCase) CreatePermission(ctx context.Context, req *request.ReqPermission) error {
	return processWithTx(ctx, uc.db, func(ctx context.Context) error {
		permission := models.Permissions{
			Code: req.Code,
			Name: req.Name,
		}
		return uc.permissionRepo.Create(ctx, &permission)
	})
}

func (uc *permissionUseCase) GetPermissionByID(ctx context.Context, id int64) (models.Permissions, error) {
	return uc.permissionRepo.GetPermissionsByID(ctx, id)
}

func (uc *permissionUseCase) GetListPermission(ctx context.Context) ([]models.Permissions, error) {
	return uc.permissionRepo.GetListPermissions(ctx)
}

func (uc *permissionUseCase) UpdatePermissionByID(ctx context.Context, req *request.ReqPermissionUpdate) (models.Permissions, error) {
	err := req.ValidateUpdatedAt()
	if err != nil {
		return models.Permissions{}, err
	}

	var updatedPermission models.Permissions
	err = processWithTx(ctx, uc.db, func(ctx context.Context) error {
		permission := models.Permissions{
			ID:   req.ID,
			Code: req.Code,
			Name: req.Name,
		}
		updatedPermission, err = uc.permissionRepo.UpdatePermissionsByID(ctx, req.ID, req.UpdatedAt, permission)
		return err
	})
	return updatedPermission, err
}

func (uc *permissionUseCase) DeletePermissionByID(ctx context.Context, id int64, updatedAt time.Time) error {
	return processWithTx(ctx, uc.db, func(ctx context.Context) error {
		return uc.permissionRepo.DeletePermissionsByID(ctx, id, updatedAt)
	})
}
