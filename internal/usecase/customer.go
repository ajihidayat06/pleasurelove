package usecase

import (
	"context"
	"time"

	"pleasurelove/internal/dto/request"
	"pleasurelove/internal/dto/response"
	"pleasurelove/internal/models"
	"pleasurelove/internal/repo"
	"pleasurelove/internal/utils"
	"pleasurelove/internal/utils/errorutils"
	"pleasurelove/pkg/logger"

	"gorm.io/gorm"
)

type CustomerUseCase interface {
	CreateCustomer(ctx context.Context, req *request.ReqCustomer) error
	GetCustomerByID(ctx context.Context, id int64) (response.CustomerResponse, error)
	GetListCustomer(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.CustomerResponse], error)
	UpdateCustomerByID(ctx context.Context, req *request.ReqCustomerUpdate) (response.CustomerResponse, error)
	DeleteCustomerByID(ctx context.Context, id int64, reqData request.AbstractRequest) error
}

type customerUseCase struct {
	db           *gorm.DB
	customerRepo repo.CustomerRepository
}

func NewCustomerUseCase(db *gorm.DB, customerRepo repo.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		db:           db,
		customerRepo: customerRepo,
	}
}

// note: setaip pelanggan mau checkout akan dibuatkan row cutomer,
// jika sudah ada maka tidak perlu dibuatkan lagi, cukup ambil dari database
// jika tidak ada maka buatkan row customer baru
// jika guest maka buatkan row customer baru dan guest token dengan user id 0
func (uc *customerUseCase) CreateCustomer(ctx context.Context, req *request.ReqCustomer) error {
	err := req.ValidateRequestCreate()
	if err != nil {
		logger.Error(ctx, "Failed to validate request", err)
		return err
	}

	// cek user login
	var (
		isGuest    bool
		geustToken string
	)
	userLogin, _ := utils.GetUserIDFromCtx(ctx)
	if userLogin == 0 {
		logger.Error(ctx, "User not logged in", errorutils.ErrDataNotFound)
		isGuest = true
	}

	if isGuest {
		geustToken, err = utils.GenerateGuestToken(req.Email, req.Phone, "web")
		if err != nil {
			logger.Error(ctx, "Failed to generate guest token", err)
			return errorutils.ErrGenerateGuestToken
		}
	}

	customer := models.Customer{
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		UserID:     userLogin,
		IsGuest:    isGuest,
		GuestToken: geustToken,
	}

	return processWithTx(ctx, uc.db, func(ctx context.Context) error {
		err := uc.customerRepo.Create(ctx, &customer)
		if err != nil {
			return errorutils.HandleRepoErrorWrite(ctx, err, repo.GetContraintErrMessage(uc.customerRepo))
		}
		return nil
	})
}

func (uc *customerUseCase) GetCustomerByID(ctx context.Context, id int64) (response.CustomerResponse, error) {
	customerDb, err := uc.customerRepo.GetCustomerByID(ctx, id)
	if err != nil {
		return response.CustomerResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	return response.SetCustomerResponse(customerDb), nil
}

func (uc *customerUseCase) GetListCustomer(ctx context.Context, listStruct *models.GetListStruct) (response.ListResponse[response.CustomerResponse], error) {
	customerDb, count, err := uc.customerRepo.GetListCustomer(ctx, listStruct)
	if err != nil {
		return response.ListResponse[response.CustomerResponse]{}, errorutils.HandleRepoError(ctx, err)
	}

	listResponse := response.MapToListResponse(response.SetResponseListCustomer(customerDb), count, listStruct, repo.GetFilterAvailableFromRepo(uc.customerRepo))
	return listResponse, nil
}

func (uc *customerUseCase) UpdateCustomerByID(ctx context.Context, req *request.ReqCustomerUpdate) (response.CustomerResponse, error) {
	err := req.ValidateUpdatedAt()
	if err != nil {
		return response.CustomerResponse{}, err
	}

	err = req.ValidateRequestUpdate()
	if err != nil {
		return response.CustomerResponse{}, err
	}

	customerDb, err := uc.customerRepo.GetCustomerByID(ctx, req.ID)
	if err != nil {
		return response.CustomerResponse{}, errorutils.HandleRepoError(ctx, err)
	}

	if !utils.ValidateUpdatedAtRequest(req.UpdatedAt, customerDb.UpdatedAt) {
		return response.CustomerResponse{}, errorutils.ErrDataDataUpdated
	}

	// userLogin, err := utils.GetUserIDFromCtx(ctx)
	// if err != nil {
	// 	logger.Error(ctx, "Failed to get user id from context", err)
	// 	return response.CustomerResponse{}, errorutils.ErrDataNotFound
	// }

	customer := models.Customer{
		ID:         req.ID,
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		UserID:     customerDb.UserID,
		IsGuest:    customerDb.IsGuest,
		GuestToken: customerDb.GuestToken,
		CreatedAt:  customerDb.CreatedAt,
		CreatedBy:  customerDb.CreatedBy,
		UpdatedAt:  time.Now(),
	}

	var (
		res models.Customer
	)

	err = processWithTx(ctx, uc.db, func(ctx context.Context) error {
		res, err = uc.customerRepo.UpdateCustomerByID(ctx, *req, customer)
		if err != nil {
			return errorutils.HandleRepoErrorWrite(ctx, err, repo.GetContraintErrMessage(uc.customerRepo))
		}
		return nil
	})

	if err != nil {
		return response.CustomerResponse{}, err
	}

	return response.SetCustomerResponse(res), nil
}

func (uc *customerUseCase) DeleteCustomerByID(ctx context.Context, id int64, reqData request.AbstractRequest) error {
	err := reqData.ValidateUpdatedAt()
	if err != nil {
		return err
	}

	customerDb, err := uc.customerRepo.GetCustomerByID(ctx, id)
	if err != nil {
		return errorutils.HandleRepoError(ctx, err)
	}

	if !utils.ValidateUpdatedAtRequest(reqData.UpdatedAt, customerDb.UpdatedAt) {
		return errorutils.ErrDataDataUpdated
	}

	return processWithTx(ctx, uc.db, func(ctx context.Context) error {
		err := uc.customerRepo.DeleteCustomerByID(ctx, id, reqData.UpdatedAt)
		if err != nil {
			return errorutils.HandleRepoError(ctx, err)
		}
		return nil
	})
}
