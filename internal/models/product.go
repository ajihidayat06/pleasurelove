package models

import "time"

type Product struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Barcode     string `json:"barcode"`
	Description string `json:"description"`

	Brand     string  `json:"brand"`
	Unit      string  `json:"unit"`
	Price     float64 `json:"price"`
	CostPrice float64 `json:"cost_price"`
	Discount  float64 `json:"discount"`

	IsActive  bool `json:"is_active"`
	HasVarian bool `json:"has_varian"`

	CreatedBy       int64              `json:"created_by"`
	UpdatedBy       int64              `json:"updated_by"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	ProductCategory *[]ProductCategory `json:"priduct_category" gorm:"foreignKey:ProductID"`
}

// TableName menentukan nama tabel custom di database
func (Product) TableName() string {
	return "product"
}
