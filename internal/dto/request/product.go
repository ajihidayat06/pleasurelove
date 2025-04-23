package request

import (
	"fmt"
	"pleasurelove/internal/utils"
)

type ReqProduct struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Code        string  `json:"code" validate:"required"`
	Barcode     string  `json:"barcode"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`      // asumsi brand berupa nama brand
	Unit        string  `json:"unit"`       // asumsi unit berupa nama satuan
	Price       float64 `json:"price"`      // harga jual
	CostPrice   float64 `json:"cost_price"` // harga modal
	Discount    float64 `json:"discount"`   // persen diskon, misal 10.5
	IsActive    bool    `json:"is_active"`
	HasVarian   bool    `json:"has_varian"`
	CategoryID  []int64 `json:"category_id"`
}

var ReqProductErrorMessage = map[string]string{
	"Code": "code required",
}

func (r *ReqProduct) ValidateRequestCreate() error {
	err := utils.ValidateCode(r.Code)
	if err != nil {
		return err
	}

	if r.Price < 0 || r.Price > 9999999999.99 {
		return fmt.Errorf("harga jual harus antara 0 - 9999999999.99")
	}

	if r.CostPrice < 0 || r.CostPrice > 9999999999.99 {
		return fmt.Errorf("harga modal harus antara 0 - 9999999999.99")
	}

	if r.Discount < 0 || r.Discount > 100 {
		return fmt.Errorf("diskon harus antara 0 - 100 persen")
	}

	r.Price = utils.RoundTo2Digits(r.Price)
	r.CostPrice = utils.RoundTo2Digits(r.CostPrice)
	r.Discount = utils.RoundTo2Digits(r.Discount)

	return nil
}

type ReqProductUpdate struct {
	ID int64 `json:"id" validate:"required"`
	// Name        string  `json:"name" validate:"required"`
	// Code        string  `json:"code"`
	// Barcode     string  `json:"barcode"`
	// Description string  `json:"description"`
	// Brand       string  `json:"brand"`      // asumsi brand berupa nama brand
	// Unit        string  `json:"unit"`       // asumsi unit berupa nama satuan
	// Price       float64 `json:"price"`      // harga jual
	// CostPrice   float64 `json:"cost_price"` // harga modal
	// Discount    float64 `json:"discount"`   // persen diskon, misal 10.5
	// IsActive    bool    `json:"is_active"`
	// HasVarian   bool    `json:"has_varian"`
	// CategoryID  []int64 `json:"category_id"`
	ReqProduct
	AbstractRequest
}

var ReqProductUpdateErrorMessage = map[string]string{
	"ID":         "id required",
	"Code":       "code required",
	"UpdateddAtStr": "updated_at required",
}

// func (r *ReqProductUpdate) ValidateRequestUpdate() error {
// 	err := utils.ValidateCode(r.Code)
// 	if err != nil {
// 		return err
// 	}

// 	if r.Price < 0 || r.Price > 9999999999.99 {
// 		return fmt.Errorf("harga jual harus antara 0 - 9999999999.99")
// 	}

// 	if r.CostPrice < 0 || r.CostPrice > 9999999999.99 {
// 		return fmt.Errorf("harga modal harus antara 0 - 9999999999.99")
// 	}

// 	if r.Discount < 0 || r.Discount > 100 {
// 		return fmt.Errorf("diskon harus antara 0 - 100 persen")
// 	}

// 	r.Price = utils.RoundTo2Digits(r.Price)
// 	r.CostPrice = utils.RoundTo2Digits(r.CostPrice)
// 	r.Discount = utils.RoundTo2Digits(r.Discount)

// 	return nil
// }
