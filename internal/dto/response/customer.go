package response

import (
	"pleasurelove/internal/models"
	"time"
)

type CustomerResponse struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	UserID     int64     `json:"user_id"`
	IsGuest    bool      `json:"is_guest"`
	GuestToken string    `json:"guest_token"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int64     `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int64     `json:"updated_by"`
}

func SetCustomerResponse(customer models.Customer) CustomerResponse {
	return CustomerResponse{
		ID:         customer.ID,
		Name:       customer.Name,
		Email:      customer.Email,
		Phone:      customer.Phone,
		UserID:     int64(customer.UserID),
		IsGuest:    customer.IsGuest,
		GuestToken: customer.GuestToken,
		CreatedAt:  customer.CreatedAt,
		CreatedBy:  int64(customer.CreatedBy),
		UpdatedAt:  customer.UpdatedAt,
		UpdatedBy:  int64(customer.UpdatedBy),
	}
}

func SetResponseListCustomer(customers []models.Customer) []CustomerResponse {
	var customerResponses []CustomerResponse
	for _, c := range customers {
		customerResponses = append(customerResponses, SetCustomerResponse(c))
	}
	return customerResponses
}
