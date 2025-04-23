package request

import (
	"errors"
	"strings"
	"time"
)

type AbstractRequest struct {
	UpdatedAt    time.Time
	UpdatedAtStr string `json:"updated_at"`
	CreateddAt   time.Time
	CreatedAtStr string `json:"created_at"`
}

func (a *AbstractRequest) ValidateUpdatedAt() error {
	if strings.TrimSpace(a.UpdatedAtStr) == "" {
		return errors.New("updated_at tidak boleh kosong")
	}

	parsedTime, err := time.Parse(time.RFC3339, a.UpdatedAtStr)
	if err != nil {
		return errors.New("format tanggal 'updated_at' tidak valid, gunakan format RFC3339 (contoh: 2025-04-20T15:04:05Z)")
	}

	a.UpdatedAt = parsedTime

	if a.UpdatedAt.IsZero() {
		return errors.New("updated_at tidak boleh kosong")
	}

	return nil
}
