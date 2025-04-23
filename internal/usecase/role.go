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
	"strings"
	"time"

	"gorm.io/gorm"
)

type RoleUseCase interface {
	CreateRole(ctx context.Context, req *request.ReqRoles) error
	GetRoleByID(ctx context.Context, id int64) (response.RolesResponse, error)
	GetListRole(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.RolesResponse], error)
	UpdateRoleByID(ctx context.Context, req *request.ReqRoleUpdate) (response.RolesResponse, error)
	DeleteRoleByID(ctx context.Context, id int64, reqData request.AbstractRequest) error
}

type roleUseCase struct {
	db                  *gorm.DB
	roleRepo            repo.RoleRepository
	permissionsRepo     repo.PermissionsRepository
	rolePermissionsRepo repo.RolePermissionsRepository
}

func NewRoleUseCase(
	db *gorm.DB,
	roleRepo repo.RoleRepository,
	permissionsRepo repo.PermissionsRepository,
	rolePermissionsRepo repo.RolePermissionsRepository,
) RoleUseCase {
	return &roleUseCase{
		db:                  db,
		roleRepo:            roleRepo,
		permissionsRepo:     permissionsRepo,
		rolePermissionsRepo: rolePermissionsRepo,
	}
}

func (uc *roleUseCase) CreateRole(ctx context.Context, req *request.ReqRoles) error {
	userLogin, err := utils.GetUserIDFromCtx(ctx)
	if err != nil {
		return errorutils.HandleCustomError(ctx, err, errorutils.ErrMessageUserNotLogin, constanta.FieldUserID)
	}

	if len(req.RolePermissions) == 0 {
		return errorutils.HandleCustomError(ctx, nil, errorutils.ErrMessaageDataRequired, constanta.FieldPermissions)
	}

	// get role by code
	roleDB, err := uc.roleRepo.GetRoleByCode(ctx, req.Code)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errorutils.HandleRepoError(ctx, err)
	}

	if strings.EqualFold(roleDB.Code, req.Code) {
		return errorutils.HandleCustomError(ctx, nil, errorutils.ErrMessaageDataAlreadyExists, constanta.FieldCode)
	}

	role := models.Roles{
		Code:      req.Code,
		Name:      req.Name,
		CreatedBy: userLogin,
		UpdatedBy: userLogin,
	}

	return processWithTx(ctx, uc.db, func(ctx context.Context) error {
		err := uc.roleRepo.Create(ctx, &role)
		if err != nil {
			logger.Error(ctx, "Failed to create role", err)
			return errorutils.HandleRepoError(ctx, err)
		}

		var rolePermissions []models.RolePermissions // check if role permissions is empty
		for _, v := range req.RolePermissions {
			permissionsDB, err := uc.permissionsRepo.GetPermissionsByID(ctx, v.PermissionID)
			if err != nil {
				return errorutils.HandleRepoError(ctx, err)
			}

			if permissionsDB.ID == 0 {
				return errorutils.HandleCustomError(ctx, nil, errorutils.ErrMessageDataNotFound, constanta.FieldPermissions)
			}

			rolePermissions = append(rolePermissions, models.RolePermissions{
				RoleID:        role.ID,
				PermissionsID: permissionsDB.ID,
				AccessScope:   v.Scope,
				CreatedBy:     userLogin,
				UpdatedBy:     userLogin,
			})
		}

		rolePermissions, err = uc.rolePermissionsRepo.Create(ctx, rolePermissions)
		if err != nil {
			return errorutils.HandleRepoError(ctx, err)
		}

		role.RolePermissions = &rolePermissions

		return nil
	})
}

func (uc *roleUseCase) GetRoleByID(ctx context.Context, id int64) (response.RolesResponse, error) {
	userDb, err := uc.roleRepo.GetRoleByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "Failed to get role by id", err)
		return response.RolesResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	return response.SetRoleDetailResponse(userDb), nil
}

func (uc *roleUseCase) GetListRole(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.RolesResponse], error) {
	rolesDb, count, err := uc.roleRepo.GetListRole(ctx, listStruct)
	if err != nil {
		logger.Error(ctx, "Failed to get list user", err)
		return response.ListResponse[response.RolesResponse]{}, errorutils.HandleRepoError(ctx, err)
	}

	listResponse := response.MapToListResponse(response.SetListResponseRole(rolesDb), count, listStruct, repo.GetFilterAvailableFromRepo(uc.roleRepo))
	return listResponse, nil
}

func (uc *roleUseCase) UpdateRoleByID(ctx context.Context, req *request.ReqRoleUpdate) (response.RolesResponse, error) {
	err := req.ValidateUpdatedAt()
	if err != nil {
		return response.RolesResponse{}, err
	}

	userLogin, err := utils.GetUserIDFromCtx(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to get user id from context", err)
		return response.RolesResponse{}, errorutils.ErrDataNotFound
	}

	roleDb, err := uc.roleRepo.GetRoleByCode(ctx, req.Code)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(ctx, "Failed to get role by code", err)
		return response.RolesResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	if roleDb.ID != 0 && roleDb.ID != req.ID {
		return response.RolesResponse{}, errorutils.HandleCustomError(ctx, nil, errorutils.ErrMessaageDataAlreadyExists, constanta.FieldCode)
	}

	if len(req.RolePermissions) == 0 {
		return response.RolesResponse{}, errorutils.HandleCustomError(ctx, nil, errorutils.ErrMessaageDataRequired, constanta.FieldPermissions)
	}

	var (
		permissionsID   []int64
		rolePermissions = []models.RolePermissions{}
	)
	for _, v := range req.RolePermissions {
		permissionsID = append(permissionsID, v.PermissionID)
	}

	listPermissions, err := uc.permissionsRepo.GetPermissionsByListID(ctx, permissionsID)
	if err != nil {
		logger.Error(ctx, "Failed to get permissions by list id", err)
		return response.RolesResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	if len(listPermissions) == 0 {
		return response.RolesResponse{}, errorutils.HandleCustomError(ctx, nil, errorutils.ErrMessageDataNotFound, constanta.FieldPermissions)
	}

	for _, v := range req.RolePermissions {
		createdBy := userLogin
		createdAt := time.Now()
		for _, perm := range listPermissions {
			if v.ID == perm.ID {
				createdBy = perm.CreatedBy
				createdAt = perm.CreatedAt
			}

		}
		rolePermissions = append(rolePermissions, models.RolePermissions{
			ID:            v.ID,
			RoleID:        req.ID,
			PermissionsID: v.PermissionID,
			AccessScope:   v.Scope,
			UpdatedBy:     userLogin,
			UpdatedAt:     time.Now(),
			CreatedAt:     createdAt,
			CreatedBy:     createdBy,
		})
	}

	role := models.Roles{
		ID:        roleDb.ID,
		Code:      req.Code,
		Name:      req.Name,
		CreatedAt: roleDb.CreatedAt,
		CreatedBy: roleDb.CreatedBy,
		UpdatedAt: time.Now(),
		UpdatedBy: userLogin,
	}

	var updatedRole models.Roles
	err = processWithTx(ctx, uc.db, func(ctx context.Context) error {
		updatedRole, err = uc.roleRepo.UpdateRoleByID(ctx, req.ID, req.UpdatedAt, role)
		if err != nil {
			logger.Error(ctx, "Failed to update role", err)
			return errorutils.HandleRepoError(ctx, err)
		}

		// delete role permissions And insert again
		err = uc.rolePermissionsRepo.DeleteRolePermissionsByRoleID(ctx, req.ID)
		if err != nil {
			logger.Error(ctx, "Failed to delete role permissions", err)
			return errorutils.HandleRepoError(ctx, err)
		}

		rolePermissions, err = uc.rolePermissionsRepo.UpdateRolePermissionsBulk(ctx, rolePermissions)
		if err != nil {
			logger.Error(ctx, "Failed to update role permissions", err)
			return errorutils.HandleRepoError(ctx, err)
		}

		updatedRole.RolePermissions = &rolePermissions

		return nil
	})

	return response.SetRoleDetailResponse(updatedRole), err
}

func (uc *roleUseCase) DeleteRoleByID(ctx context.Context, id int64, reqData request.AbstractRequest) error {
	err := reqData.ValidateUpdatedAt()
	if err != nil {
		return err
	}

	roleDb, err := uc.roleRepo.GetRoleByID(ctx, id)
	if err != nil {
		return errorutils.HandleRepoError(ctx, err)
	}

	if roleDb.ID == 0 {
		return errorutils.ErrDataNotFound
	}

	if !utils.ValidateUpdatedAtRequest(reqData.UpdatedAt, roleDb.UpdatedAt) {
		return errorutils.ErrDataDataUpdated
	}

	return processWithTx(ctx, uc.db, func(ctx context.Context) error {
		uc.rolePermissionsRepo.DeleteRolePermissionsByRoleID(ctx, id)
		if err != nil {
			logger.Error(ctx, "Failed to delete role permissions", err)
			return errorutils.HandleRepoError(ctx, err)
		}

		err := uc.roleRepo.DeleteRoleByID(ctx, id, reqData.UpdatedAt)
		if err != nil {
			logger.Error(ctx, "Failed to delete user", err)
			return errorutils.HandleRepoError(ctx, err)
		}
		return nil
	})
}
