package middleware

import (
	"encoding/json"
	"fmt"
	"pleasurelove/internal/utils"
	"pleasurelove/pkg/logger"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	ctx := utils.GetContext(c)
	start := time.Now()

	// Simpan request body untuk logging
	requestBody := string(c.Body())

	// Hindari mencatat informasi sensitif
	requestBody = sanitizeSensitiveData(requestBody)

	// Jalankan handler berikutnya
	err := c.Next()

	// Hitung durasi request
	duration := time.Since(start)

	// Siapkan data log
	logData := map[string]interface{}{
		"timestamp":    time.Now().Format(time.RFC3339),
		"method":       c.Method(),
		"path":         c.Path(),
		"query_params": string(c.Request().URI().QueryString()),
		"request_body": requestBody,
		"status":       c.Response().StatusCode(),
		"duration":     duration.String(),
	}

	// Log berdasarkan status
	if err != nil {
		// Jika terjadi error, log sebagai ERROR
		logData["error"] = sanitizeError(err).Error()
		logger.Error(ctx, fmt.Sprintf("Request Error: %v | Data: %v", err, logData), err)
	} else {
		// Jika tidak ada error, log sebagai INFO
		logger.Info(ctx, "Request Log", logData)
	}

	return err
}

// sanitizeSensitiveData mengganti informasi sensitif dalam request body dengan [FILTERED]
func sanitizeSensitiveData(body string) string {
	// Coba parsing body sebagai JSON
	var parsedBody map[string]interface{}
	if err := json.Unmarshal([]byte(body), &parsedBody); err != nil {
		// Jika gagal parsing, coba untuk menggunakan sanitizeString
		return sanitizeString(body, []string{"password", "token", "api_key", "secret"})
	}

	// Daftar key sensitif
	sensitiveFields := []string{"password", "token", "api_key", "secret"}

	// Iterasi melalui key-value dan filter value sensitif
	for key, value := range parsedBody {
		if strVal, ok := value.(string); ok {
			// Jika key termasuk sensitive, langsung filter
			if containsInsensitive(sensitiveFields, key) {
				parsedBody[key] = "[FILTERED]"
			} else {
				parsedBody[key] = sanitizeString(strVal, sensitiveFields)
			}
		}
	}

	// Kembalikan body yang sudah difilter sebagai JSON string
	filteredBody, err := json.Marshal(parsedBody)
	if err != nil {
		// Jika gagal mengubah kembali ke JSON, kembalikan body asli
		return body
	}
	return string(filteredBody)
}

// sanitizeError memastikan error yang dicatat tidak mengandung informasi sensitif
func sanitizeError(err error) error {
	sensitiveKeywords := []string{"password", "token", "api_key", "secret"}
	sanitizedMessage := err.Error()
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(strings.ToLower(sanitizedMessage), keyword) {
			sanitizedMessage = strings.ReplaceAll(sanitizedMessage, keyword, "[FILTERED]")
		}
	}
	return fiber.NewError(fiber.StatusInternalServerError, sanitizedMessage)
}

func sanitizeString(input string, sensitiveKeywords []string) string {
	for _, keyword := range sensitiveKeywords {
		re := regexp.MustCompile(`(?i)` + keyword + `\s*[:=]\s*[^&\s]+`)
		input = re.ReplaceAllString(input, keyword+"=[FILTERED]")
	}
	return input
}

func containsInsensitive(list []string, target string) bool {
	for _, item := range list {
		if strings.EqualFold(item, target) {
			return true
		}
	}
	return false
}
