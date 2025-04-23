package response

import (
	"pleasurelove/internal/models"
	"pleasurelove/internal/utils"
	"time"
)

type ProductResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Price     float64   `json:"price"`
	CostPrice float64   `json:"cost_price"`
	Discount  float64   `json:"discount"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int64     `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int64     `json:"updated_by"`
}

func SetProductResponse(product models.Product) ProductResponse {
	return ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Code:      product.Code,
		Price:     utils.RoundTo2Digits(product.Price),
		CostPrice: utils.RoundTo2Digits(product.CostPrice),
		Discount:  utils.RoundTo2Digits(product.Discount),
		IsActive:  product.IsActive,
		CreatedAt: product.CreatedAt,
		CreatedBy: product.CreatedBy,
		UpdatedAt: product.UpdatedAt,
		UpdatedBy: product.UpdatedBy,
	}
}

func SetResponseListProduct(products []models.Product) []ProductResponse {
	var responses []ProductResponse
	for _, product := range products {
		responses = append(responses, SetProductResponse(product))
	}
	return responses
}

type DetailProductResponse struct {
	ID              int64                     `json:"id"`
	Name            string                    `json:"name"`
	Code            string                    `json:"code"`
	Barcode         string                    `json:"barcode"`
	Description     string                    `json:"description"`
	Brand           string                    `json:"brand"`
	Unit            string                    `json:"unit"`
	Price           float64                   `json:"price"`
	CostPrice       float64                   `json:"cost_price"`
	Discount        float64                   `json:"discount"`
	IsActive        bool                      `json:"is_active"`
	HasVarian       bool                      `json:"has_varian"`
	CreatedAt       time.Time                 `json:"created_at"`
	CreatedBy       int64                     `json:"created_by"`
	UpdatedAt       time.Time                 `json:"updated_at"`
	UpdatedBy       int64                     `json:"updated_by"`
	ProductCategory []ProductCategoryResponse `json:"product_category"`
}

func SetDetailProductResponse(product models.Product) DetailProductResponse {
	var productcategory []ProductCategoryResponse
	for _, pc := range *product.ProductCategory {
		productcategory = append(productcategory, SetProductCategoryResponse(pc))
	}

	return DetailProductResponse{
		ID:              product.ID,
		Name:            product.Name,
		Code:            product.Code,
		Barcode:         product.Barcode,
		Description:     product.Description,
		Brand:           product.Brand,
		Unit:            product.Unit,
		Price:           utils.RoundTo2Digits(product.Price),
		CostPrice:       utils.RoundTo2Digits(product.CostPrice),
		Discount:        utils.RoundTo2Digits(product.Discount),
		IsActive:        product.IsActive,
		HasVarian:       product.HasVarian,
		CreatedAt:       product.CreatedAt,
		CreatedBy:       product.CreatedBy,
		UpdatedAt:       product.UpdatedAt,
		UpdatedBy:       product.UpdatedBy,
		ProductCategory: productcategory,
	}
}
