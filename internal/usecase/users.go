package usecase

import (
	"context"
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

type UserUseCase interface {
	Register(ctx context.Context, reqUser *request.ReqUser) error
	Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error)
	GetUserByID(ctx context.Context, user int64) (response.UserResponse, error)
	CreateUserDashboard(ctx context.Context, user *request.ReqUser) error
	GetListUser(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.UserResponse], error)
	UpdateUserByID(ctx context.Context, user *request.ReqUserUpdate) (response.UserResponse, error)
	DeleteUserByID(ctx context.Context, id int64, reqData request.AbstractRequest) error
}

type userUseCase struct {
	db       *gorm.DB
	UserRepo repo.UserRepository
	RoleRepo repo.RoleRepository
}

func NewUserUseCase(db *gorm.DB, userRepo repo.UserRepository, roleRepo repo.RoleRepository) UserUseCase {
	return &userUseCase{
		db:       db,
		UserRepo: userRepo,
		RoleRepo: roleRepo,
	}
}

func (u *userUseCase) Register(ctx context.Context, reqUser *request.ReqUser) error {
	user := models.User{}
	return processWithTx(ctx, u.db, func(ctx context.Context) error {
		err := u.UserRepo.Create(ctx, &user)
		if err != nil {
			logger.Error(ctx, "Failed to create user", err)
			return err
		}
		return nil
	})
}

func (u userUseCase) Login(ctx context.Context, req *request.ReqLogin) (models.UserLogin, error) {
	// TODO: get user by (username or email) and password
	user := models.UserLogin{}
	return user, nil
}

func (u *userUseCase) GetUserByID(ctx context.Context, id int64) (response.UserResponse, error) {
	userDb, err := u.UserRepo.GetUserByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "Failed to get user by id", err)
		return response.UserResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	return response.SetUserDetailResponse(userDb), nil
}

func (u *userUseCase) CreateUserDashboard(ctx context.Context, reqUser *request.ReqUser) error {
	err := reqUser.ValidateRequestCreate()
	if err != nil {
		return err
	}

	userLogin, err := utils.GetUserIDFromCtx(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to get user id from context", err)
		return errorutils.ErrDataNotFound
	}

	user := models.User{
		Name:      reqUser.Name,
		Email:     reqUser.Email,
		Username:  reqUser.Username,
		Password:  reqUser.Password,
		RoleID:    reqUser.RoleID,
		CreatedAt: time.Now(),
		CreatedBy: userLogin,
		UpdatedAt: time.Now(),
		UpdatedBy: userLogin,
	}

	var (
		roleDb models.Roles
	)

	switch {
	case reqUser.RoleID != 0:
		roleDb, err = u.getDataRole(ctx, reqUser.RoleID)
		if err != nil {
			return err
		}

		user.RoleID = roleDb.ID
	case reqUser.RoleID == 0 && (strings.TrimSpace(reqUser.Roles.Name) != "" || strings.TrimSpace(reqUser.Roles.Code) != ""):
		roleCode := strings.ToLower(reqUser.Roles.Code)
		if roleCode == constanta.RoleCodeAdmin {
			roleCode = constanta.RoleCodeAdmin
		} else {
			roleCode = reqUser.Roles.Code
		}

		roles := models.Roles{
			Name:      reqUser.Roles.Name,
			Code:      roleCode,
			CreatedBy: userLogin,
			UpdatedBy: userLogin,
		}
		user.Roles = &roles
	}

	return processWithTx(ctx, u.db, func(ctx context.Context) error {
		err := u.UserRepo.Create(ctx, &user)
		if err != nil {
			logger.Error(ctx, "Failed to create user dashboard", err)
			return errorutils.HandleRepoError(ctx, err)
		}
		return nil
	})
}

func (u *userUseCase) GetListUser(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.UserResponse], error) {
	userDb, count, err := u.UserRepo.GetListUser(ctx, listStruct)
	if err != nil {
		logger.Error(ctx, "Failed to get list user", err)
		return response.ListResponse[response.UserResponse]{}, errorutils.HandleRepoError(ctx, err)
	}

	listResponse := response.MapToListResponse(response.SetResponseListUser(userDb), count, listStruct, repo.GetFilterAvailableFromRepo(u.UserRepo))
	return listResponse, nil
}

func (u *userUseCase) UpdateUserByID(ctx context.Context, reqData *request.ReqUserUpdate) (response.UserResponse, error) {
	err := reqData.ValidateRequestUpdate()
	if err != nil {
		return response.UserResponse{}, err
	}

	userDb, err := u.UserRepo.GetUserByID(ctx, reqData.ID)
	if err != nil {
		return response.UserResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	if userDb.ID == 0 {
		return response.UserResponse{}, errorutils.ErrDataNotFound
	}

	if !utils.ValidateUpdatedAtRequest(reqData.UpdatedAt, userDb.UpdatedAt) {
		return response.UserResponse{}, errorutils.ErrDataDataUpdated
	}

	var roleDb models.Roles
	if reqData.RoleID != 0 {
		roleDb, err = u.getDataRole(ctx, reqData.RoleID)
		if err != nil {
			return response.UserResponse{}, errorutils.HandleRepoError(ctx, err)
		}
	}

	userLogin, err := utils.GetUserIDFromCtx(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to get user id from context", err)
		return response.UserResponse{}, errorutils.ErrDataNotFound
	}

	userUpdate := models.User{
		ID:        userDb.ID,
		Name:      reqData.Name,
		Email:     reqData.Email,
		Username:  reqData.Username,
		Password:  reqData.Password,
		RoleID:    reqData.RoleID,
		Roles:     &roleDb,
		CreatedAt: userDb.CreatedAt,
		CreatedBy: userDb.CreatedBy,
		UpdatedAt: time.Now(),
		UpdatedBy: userLogin,
	}

	var (
		res models.User
	)
	err = processWithTx(ctx, u.db, func(ctx context.Context) error {
		res, err = u.UserRepo.UpdateUserByID(ctx, *reqData, userUpdate)
		if err != nil {
			logger.Error(ctx, "Failed to update user", err)
			return err
		}

		return nil
	})
	if err != nil {
		return response.UserResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	return response.SetUserDetailResponse(res), nil
}

func (u *userUseCase) DeleteUserByID(ctx context.Context, id int64, reqData request.AbstractRequest) error {
	err := reqData.ValidateUpdatedAt()
	if err != nil {
		return err
	}

	userDb, err := u.UserRepo.GetUserByID(ctx, id)
	if err != nil {
		return errorutils.HandleRepoError(ctx, err)
	}

	if userDb.ID == 0 {
		return errorutils.ErrDataNotFound
	}

	if !utils.ValidateUpdatedAtRequest(reqData.UpdatedAt, userDb.UpdatedAt) {
		return errorutils.ErrDataDataUpdated
	}

	return processWithTx(ctx, u.db, func(ctx context.Context) error {
		err := u.UserRepo.DeleteUserByID(ctx, id, reqData.UpdatedAt)
		if err != nil {
			return errorutils.HandleRepoError(ctx, err)
		}
		return nil
	})
}

func (u *userUseCase) getDataRole(ctx context.Context, roleID int64) (res models.Roles, err error) {
	if roleID != 0 {
		res, err = u.RoleRepo.GetRoleByID(ctx, roleID)
		if err != nil {
			logger.Error(ctx, "Failed to get role by id", err)
			return models.Roles{}, errorutils.HandleRepoError(ctx, err)
		}
	}
	return
}
