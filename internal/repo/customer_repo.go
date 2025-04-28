package repo

import (
	"context"
	"pleasurelove/internal/dto/request"
	"pleasurelove/internal/models"
	"time"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *models.Customer) error
	GetCustomerByID(ctx context.Context, id int64) (models.Customer, error)
	GetCustomerByUserID(ctx context.Context, userID int64) (models.Customer, error)
	GetListCustomer(ctx context.Context, listStruct *models.GetListStruct) ([]models.Customer, int64, error)
	UpdateCustomerByID(ctx context.Context, reqData request.ReqCustomerUpdate, customer models.Customer) (models.Customer, error)
	DeleteCustomerByID(ctx context.Context, id int64, updatedAt time.Time) error
}

type customerRepository struct {
	AbstractRepo
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		AbstractRepo: AbstractRepo{
			db: db,
		},
	}
}

func (r *customerRepository) Create(ctx context.Context, customer *models.Customer) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).Create(customer).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) GetCustomerByID(ctx context.Context, id int64) (models.Customer, error) {
	var customer models.Customer

	err := r.db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Where("id = ?", id).
		First(&customer).Error
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func (r *customerRepository) GetCustomerByUserID(ctx context.Context, userID int64) (models.Customer, error) {
	var customer models.Customer

	err := r.db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Where("user_id = ?", userID).
		First(&customer).Error
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func (r *customerRepository) GetListCustomer(ctx context.Context, listStruct *models.GetListStruct) ([]models.Customer, int64, error) {
	var customers []models.Customer
	var total int64

	err := r.db.WithContext(ctx).
		Model(&models.Customer{}).
		Scopes(r.withCheckScope(ctx), r.applyFilters(listStruct.Filters)).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).
		Model(&models.Customer{}).
		Scopes(r.withCheckScope(ctx), r.applyFiltersAndPaginationAndOrder(listStruct)).
		Find(&customers).Error
	if err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

func (r *customerRepository) UpdateCustomerByID(ctx context.Context, reqData request.ReqCustomerUpdate, customer models.Customer) (models.Customer, error) {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Model(&customer).
		Where("id = ? AND updated_at = ?", reqData.ID, reqData.UpdatedAt).
		Updates(&customer).Error
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func (r *customerRepository) DeleteCustomerByID(ctx context.Context, id int64, updatedAt time.Time) error {
	db := r.getDB(ctx)

	err := db.WithContext(ctx).
		Scopes(r.withCheckScope(ctx)).
		Where("id = ? AND updated_at = ?", id, updatedAt).
		Delete(&models.Customer{}).Error
	if err != nil {
		return err
	}

	return nil
}
