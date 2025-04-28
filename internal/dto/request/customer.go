package request

import "pleasurelove/internal/utils"

type ReqCustomer struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required"`
	UserID  int    `json:"user_id"`
	IsGuest bool   `json:"is_guest"`
}

var ReqCustomerErrorMessage = map[string]string{
	"Name":  "name is required",
	"Email": "valid email is required",
	"Phone": "phone number is required",
}

func (r *ReqCustomer) ValidateRequestCreate() error {
	err := utils.ValidateEmail(r.Email)
	if err != nil {
		return err
	}

	err = utils.ValidatePhone(r.Phone)
	if err != nil {
		return err
	}

	return nil
}

type ReqCustomerUpdate struct {
	ID      int64  `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required"`
	UserID  int    `json:"user_id"`
	IsGuest bool   `json:"is_guest"`
	AbstractRequest
}

var ReqCustomerUpdateErrorMessage = map[string]string{
	"ID":            "id is required",
	"Name":          "name is required",
	"Email":         "valid email is required",
	"Phone":         "phone number is required",
	"UpdateddAtStr": "updated_at is required",
}

func (r *ReqCustomerUpdate) ValidateRequestUpdate() error {
	if err := r.ValidateUpdatedAt(); err != nil {
		return err
	}

	err := utils.ValidateEmail(r.Email)
	if err != nil {
		return err
	}

	err = utils.ValidatePhone(r.Phone)
	if err != nil {
		return err
	}

	return nil
}
