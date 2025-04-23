package request

import (
	"context"
	"errors"
	"pleasurelove/pkg/logger"
)

type ReqLogin struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8,max=20"`
}

func (r *ReqLogin) ValidateRequest(ctx context.Context) error {
	// Validasi input
	if r.UsernameOrEmail == "" || r.Password == "" {
		logger.Error(ctx, "Login attempt with empty username or password", nil)
		return errors.New("Username and password are required")
	}

	return nil
}

var ReqLoginErrorMessage = map[string]string{
	"username_or_email": "invalid username or email",
	"password":          "invalid password",
}

type ReqToken struct {
	TemporaryToken string `json:"temporary_token" validate:"required"`
}

func (r *ReqToken) ValidateRequest(ctx context.Context) error {
	// Validasi input
	if r.TemporaryToken == "" {
		err := errors.New("Token are required")
		logger.Error(ctx, "temporary token nil", err)
		return err
	}

	return nil
}
