package request

import (
	"pleasurelove/internal/utils"
	"pleasurelove/internal/utils/errorutils"
)

// req user
type ReqUser struct {
	ID       int64    `json:"id"`
	Username string   `json:"username" validate:"required"`
	Name     string   `json:"name" validate:"required"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required"`
	RoleID   int64    `json:"role_id"`
	Roles    ReqRoles `json:"roles"`
}

var ReqUserErrorMessage = map[string]string{
	"name":     "name required",
	"email":    "email not valid",
	"username": "username required",
	"password": "password required",
}

func (r *ReqUser) ValidateRequestCreate() error {
	err := utils.ValidateEmail(r.Email)
	if err != nil {
		return err
	}

	err = utils.ValidateUsername(r.Username)
	if err != nil {
		return err
	}

	isValid := utils.ValidatePassword(r.Password)
	if !isValid {
		return errorutils.ErrPasswordNotValid
	}

	r.Password, err = utils.HashPassword(r.Password)
	if err != nil {
		return err
	}

	return nil
}

type ReqUserUpdate struct {
	ID       int64  `json:"id"`
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	RoleID   int64  `json:"role_id"`
	AbstractRequest
}

var ReqUserUpdateErrorMessage = map[string]string{
	"ID":            "id required",
	"Name":          "name required",
	"Email":         "email not valid",
	"Username":      "username required",
	"Password":      "password required",
	"UpdateddAtStr": "updated_at required",
}

func (r *ReqUserUpdate) ValidateRequestUpdate() error {
	if err := r.ValidateUpdatedAt(); err != nil {
		return err
	}

	err := utils.ValidateEmail(r.Email)
	if err != nil {
		return err
	}

	err = utils.ValidateUsername(r.Username)
	if err != nil {
		return err
	}

	isValid := utils.ValidatePassword(r.Password)
	if !isValid {
		return errorutils.ErrPasswordNotValid
	}

	r.Password, err = utils.HashPassword(r.Password)
	if err != nil {
		return err
	}
	return nil
}
