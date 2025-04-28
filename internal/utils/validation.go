package utils

import (
	"errors"
	"math"
	"pleasurelove/internal/constanta"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
}

func ValidateRequest(input interface{}, customMessages map[string]string) (bool, string) {
	err := Validate.Struct(input)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldErr := range validationErrors {
				fieldName := fieldErr.Field()
				if msg, exists := customMessages[fieldName]; exists {
					errorMessages = append(errorMessages, msg)
				} else {
					// Pesan default jika tidak ada mapping khusus
					errorMessages = append(errorMessages, fieldName+": "+fieldErr.Tag())
				}
			}
			return false, strings.Join(errorMessages, ";")
		}
		return false, err.Error()
	}
	return true, ""
}

func ValidateUsername(username string) error {
	usernameRegex := regexp.MustCompile(constanta.UsernameRegex)
	if !usernameRegex.MatchString(username) {
		return errors.New("format username tidak valid (hanya huruf, angka, dan underscores, dengan panjang 3-40 karakter) ")
	}
	return nil
}

func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(constanta.EmailRegex)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidateLoginInput(input string) error {
	usernameRegex := regexp.MustCompile(constanta.UsernameRegex)
	emailRegex := regexp.MustCompile(constanta.EmailRegex)

	switch {
	case emailRegex.MatchString(input):
		return nil
	case usernameRegex.MatchString(input):
		return nil
	default:
		return errors.New("invalid email or username format")
	}
}

// ValidatePassword memvalidasi password sesuai kriteria
func ValidatePassword(password string) bool {
	// Cek panjang password (minimal 8, maksimal 20)
	if len(password) < 8 || len(password) > 20 {
		return false
	}

	// Cek minimal satu huruf kecil
	lowercase := regexp.MustCompile(`[a-z]`)
	if !lowercase.MatchString(password) {
		return false
	}

	// Cek minimal satu huruf besar
	uppercase := regexp.MustCompile(`[A-Z]`)
	if !uppercase.MatchString(password) {
		return false
	}

	// Cek minimal satu angka
	number := regexp.MustCompile(`\d`)
	if !number.MatchString(password) {
		return false
	}

	// Cek minimal satu karakter spesial
	specialChar := regexp.MustCompile(`[!@#$%^&*()_+\[\]{}|;:'",.<>?/\\]`)
	return specialChar.MatchString(password)
}

func ValidateUpdatedAtRequest(request, dataDb time.Time) bool {
	return request.Equal(dataDb)
}

func ValidateCode(code string) error {
	codeRegex := regexp.MustCompile(constanta.UsernameRegex)
	if !codeRegex.MatchString(code) {
		return errors.New("format code tidak valid (hanya huruf, angka, dan underscores, dengan panjang 3-40 karakter) ")
	}
	return nil
}

func GenerateSlug(input string) string {
	slug := strings.ToLower(input)

	// Ganti semua non-alphanumeric (selain huruf dan angka) dengan dash
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}

func RoundTo2Digits(value float64) float64 {
	return math.Round(value*100) / 100
}

func ValidatePhone(phone string) error {
	phone = strings.TrimSpace(phone)

	// Regular expression: start with optional +, followed by digits only
	re := regexp.MustCompile(constanta.PhoneRegex) // min 7, max 15 digits (standar umum telepon)

	if !re.MatchString(phone) {
		return errors.New("format nomor telepon tidak valid (hanya angka, panjang 7-15 karakter)")
	}
	return nil
}
