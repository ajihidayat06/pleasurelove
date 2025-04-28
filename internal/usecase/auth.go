package usecase

import (
	"context"
	"pleasurelove/internal/dto/request"
	"pleasurelove/internal/models"
	"pleasurelove/internal/repo"
	"pleasurelove/internal/utils"
	"pleasurelove/internal/utils/errorutils"

	"gorm.io/gorm"
)

type AuthUseCase interface {
	Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	// Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	LogoutDashboard(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	// Logout(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	LoginByUserId(ctx context.Context, userID int64) (models.UserLogin, error)
}

type authUseCase struct {
	db       *gorm.DB
	UserRepo repo.UserRepository
}

func NewAuthUseCase(db *gorm.DB, userRepo repo.UserRepository) AuthUseCase {
	return &authUseCase{
		db:       db,
		UserRepo: userRepo,
	}
}

func (u *authUseCase) Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
	// Ambil user dari repository
	user, err := u.UserRepo.Login(ctx, req.UsernameOrEmail, 0)
	if err != nil {
		return models.UserLogin{}, err
	}

	// Validasi password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return models.UserLogin{}, errorutils.ErrInvalidCredentials // Ensure the error is defined in the errors package
	}

	// Mapping RolePermissions ke UserLogin
	var rolePermissions []models.RolePermissions
	for _, rp := range *user.Roles.RolePermissions {
		rolePermissions = append(rolePermissions, models.RolePermissions{
			ID:            rp.ID,
			RoleID:        rp.RoleID,
			PermissionsID: rp.PermissionsID,
			AccessScope:   rp.AccessScope,
			Permissions: &models.Permissions{
				ID:        rp.Permissions.ID,
				Code:      rp.Permissions.Code,
				Name:      rp.Permissions.Name,
				Action:    rp.Permissions.Action,
				GroupMenu: rp.Permissions.GroupMenu,
			},
		})
	}

	// Buat UserLogin
	userLogin := models.UserLogin{
		ID:              user.ID,
		RoleID:          user.RoleID,
		RoleName:        user.Roles.Name,
		RoleCode:        user.Roles.Code,
		RolePermissions: rolePermissions,
	}

	return userLogin, nil
}

// Login implements AuthUseCase.
// func (u *authUseCase) Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
// 	panic("unimplemented")
// }

// Logout implements AuthUseCase.
// func (u *authUseCase) Logout(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
// 	panic("unimplemented")
// }

// LogoutDashboard implements AuthUseCase.
func (u *authUseCase) LogoutDashboard(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
	panic("unimplemented")
}

func (u *authUseCase) LoginByUserId(ctx context.Context, userID int64) (models.UserLogin, error) {
	// Ambil user dari repository berdasarkan userID
	user, err := u.UserRepo.Login(ctx, "", userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.UserLogin{}, errorutils.ErrDataNotFound // Pastikan error ini didefinisikan di package errors
		}
		return models.UserLogin{}, err
	}

	// Mapping RolePermissions ke UserLogin
	var rolePermissions []models.RolePermissions
	for _, rp := range *user.Roles.RolePermissions {
		rolePermissions = append(rolePermissions, models.RolePermissions{
			ID:            rp.ID,
			RoleID:        rp.RoleID,
			PermissionsID: rp.PermissionsID,
			AccessScope:   rp.AccessScope,
			Permissions: &models.Permissions{
				ID:        rp.Permissions.ID,
				Code:      rp.Permissions.Code,
				Name:      rp.Permissions.Name,
				Action:    rp.Permissions.Action,
				GroupMenu: rp.Permissions.GroupMenu,
			},
		})
	}

	// Buat UserLogin
	userLogin := models.UserLogin{
		ID:              user.ID,
		RoleID:          user.RoleID,
		RoleName:        user.Roles.Name,
		RoleCode:        user.Roles.Code,
		RolePermissions: rolePermissions,
	}

	return userLogin, nil
}
