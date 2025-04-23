package errorutils

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"pleasurelove/internal/dto/response"
	"pleasurelove/internal/utils"
	"pleasurelove/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	ErrMessageForbidden             = "Anda tidak memiliki hak akses untuk mengakses menu ini."
	ErrMessageInvalidToken          = "Token tidak sesuai"
	ErrMessageExpiredToken          = "Token kadaluwarsa"
	ErrMessageInvalidOrExpiredToken = "Token tidak sesuai atau kadaluwarsa"
	ErrMessageDataNotFound          = "Data tidak ditemukan"
	ErrMessageInvalidRequestData    = "Data tidak valid"
	ErrMessageDataUpdated           = "data telah diperbarui, silakan muat ulang data terbaru"
	ErrMessaageDataAlreadyExists    = "data sudah ada"
	ErrMessaageDataRequired         = "data tidak boleh kosong"
	ErrMessageUserNotLogin          = "silahkan login terlebih dahulu"
	ErrMessageInternalServerError   = "terjadi kesalahan pada server, silahkan hubungi admin"
)

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrDataNotFound        = errors.New("data tidak ditemukan")
	ErrInternalServerError = errors.New("internal server error")
	ErrPasswordNotValid    = errors.New("password must contain at least 8 characters, including uppercase, lowercase, numbers, and special characters")
	ErrDataDataUpdated     = errors.New(ErrMessageDataUpdated)
	ErrDataAlreadyExists   = errors.New(ErrMessaageDataAlreadyExists)
)

type CustomError struct {
	Message    string
	FieldError string
	Err        error
}

func (e *CustomError) Error() string {
	if e.FieldError != "" {
		return fmt.Sprintf("%s [%s]", e.Message, e.FieldError)
	}
	return e.Message
}

func HandleCustomError(ctx context.Context, baseErr error, msg string, fieldError ...string) error {
	var fe string
	if len(fieldError) > 0 {
		fe = fieldError[0]
	}

	// Kalau baseErr nil, tetap buat error berdasarkan msg dan fieldError
	if baseErr == nil {
		baseErr = fmt.Errorf("%s", msg)
	}

	wrappedErr := &CustomError{
		Message:    msg,
		FieldError: fe,
		Err:        baseErr,
	}

	logger.LogWithCaller(ctx, msg, baseErr, 2)

	return wrappedErr
}

func HandleRepoError(ctx context.Context, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.LogWithCaller(ctx, ErrMessageDataNotFound, err, 2)
		return ErrDataNotFound
	}

	logger.LogWithCaller(ctx, ErrInternalServerError.Error(), err, 2)
	return ErrInternalServerError
}

func HandleUsecaseError(c *fiber.Ctx, err error, msg string) error {
	ctx := utils.GetContext(c)

	if errors.Is(err, ErrDataNotFound) {
		logger.LogWithCaller(ctx, ErrMessageDataNotFound, err, 2)
		return response.SetResponseNotFound(c, ErrMessageDataNotFound, err)
	}

	if errors.Is(err, ErrInternalServerError) {
		logger.LogWithCaller(ctx, ErrInternalServerError.Error(), err, 2)
		return response.SetResponseInternalServerError(c, ErrMessageInternalServerError, err)
	}

	logger.LogWithCaller(ctx, msg, err, 2)
	return response.SetResponseBadRequest(c, msg, err)
}

func HandleRepoErrorWrite(ctx context.Context, err error, constraintErr map[string]string) error {
	duplicate := "duplicate key value violates unique constraint"
	if strings.Contains(err.Error(), duplicate) {
		msg := GetMessageConstraintError(err, constraintErr)
		logger.LogWithCaller(ctx, msg, ErrDataAlreadyExists, 2)
		return HandleCustomError(ctx, err, msg)
	}

	return HandleRepoError(ctx, err)
}

func GetMessageConstraintError(err error, constraintErrorMessages map[string]string) string {
	msg := err.Error()

	start := strings.Index(msg, `"`)
	end := strings.LastIndex(msg, `"`)

	if start == -1 || end == -1 || start >= end {
		return "Data sudah ada"
	}

	constraint := msg[start+1 : end]
	if friendly, ok := constraintErrorMessages[constraint]; ok {
		return friendly
	}

	return "Data sudah ada"
}
