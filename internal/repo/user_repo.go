package repo

import (
	"context"
	"pleasurelove/internal/dto/request"
	"pleasurelove/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Login(ctx context.Context, emailOrUsername string, userID int64) (*models.User, error)
	GetUserByID(ctx context.Context, id int64) (models.User, error)
	GetListUser(ctx context.Context, listStruct *models.GetListStruct) ([]models.User, int64, error)
	UpdateUserByID(ctx context.Context, reqData request.ReqUserUpdate, user models.User) (models.User, error)
	DeleteUserByID(ctx context.Context, id int64, updatedAt time.Time) error
}

type userRepository struct {
	AbstractRepo
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		AbstractRepo: AbstractRepo{
			db: db,
		},
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Login(ctx context.Context, emailOrUsername string, userID int64) (*models.User, error) {
	var user models.User
	db := r.getDB(ctx).Preload("Roles").
		Preload("Roles.RolePermissions").
		Preload("Roles.RolePermissions.Permissions")

	if userID != 0 {
		db.Where("id = ?", userID)
	} else {
		db.Where("(email = ? OR username = ?)", emailOrUsername, emailOrUsername)
	}

	err := db.First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Preload("Roles").
		Where(" id = ? ", id).
		First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetListUser(ctx context.Context, listStruct *models.GetListStruct) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Query untuk menghitung total data tanpa Preload atau Pagination
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Scopes(r.withCheckScope(ctx), r.applyFilters(listStruct.Filters)).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Query untuk mendapatkan data dengan Pagination dan Sorting
	err = r.db.WithContext(ctx).
		Model(&models.User{}).Preload("Roles").
		Scopes(r.withCheckScope(ctx), r.applyFiltersAndPaginationAndOrder(listStruct)).
		Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) UpdateUserByID(ctx context.Context, reqData request.ReqUserUpdate, user models.User) (models.User, error) {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Model(&user).
		Where("id = ? AND updated_at = ?", reqData.ID, reqData.UpdatedAt).
		Updates(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUserByID(ctx context.Context, id int64, updatedAt time.Time) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Delete(&models.User{}).Error
	if err != nil {
		return err
	}

	return nil
}
