package models

import "time"

type Customer struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	UserID     int64     `json:"user_id"`
	IsGuest    bool      `json:"is_guest"`
	GuestToken string    `json:"guest_token"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int       `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int       `json:"updated_by"`
}

func (Customer) TableName() string {
	return "customer"
}
