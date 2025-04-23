package utils

import (
	"context"
	"errors"
	"pleasurelove/internal/constanta"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIDFromCtx(ctx context.Context) (int64, error) {
	userID := ctx.Value(constanta.AuthUserID)
	if userID == nil {
		return 0, errors.New("user_id tidak ditemukan di context atau tipe tidak sesuai")
	}
	return userID.(int64), nil
}

func GetContext(c *fiber.Ctx) context.Context {
	return c.UserContext()
}

func GenerateTraceID() string {
	return uuid.New().String() // Generate a new UUID as trace ID
}

// generateRequestID generates a unique request ID
func GenerateRequestID() string {
	return uuid.New().String() // Generate a new UUID as request ID
}
