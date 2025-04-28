package utils

import (
	"context"
	"errors"
	"pleasurelove/internal/constanta"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

type GuestClaims struct {
	GuestID string `json:"guest_id"`
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Source  string `json:"source,omitempty"`
	jwt.RegisteredClaims
}

func GenerateGuestToken(email, phone, source string) (string, error) {
	guestID := GenerateRequestID()

	claims := GuestClaims{
		GuestID: guestID,
		Email:   email,
		Phone:   phone,
		Source:  source,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // Token berlaku 7 hari
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	guestSecret := GetOSEnvGuestSecretKey()
	if guestSecret == "" {
		return "", errors.New("guest secret key not found in environment variables")
	}

	signedToken, err := token.SignedString(guestSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
