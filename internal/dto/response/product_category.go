package response

import (
	"pleasurelove/internal/models"
	"time"
)

type ProductCategoryResponse struct {
	ID             int64            `json:"id"`
	ProductID      int64            `json:"product_id"`
	CategoryID     int64            `json:"categroy_id"`
	CategoryDetail CategoryResponse `json:"category_detail"`
	CreatedAt      time.Time        `json:"created_at"`
	CreatedBy      int64            `json:"created_by"`
	UpdatedAt      time.Time        `json:"updated_at"`
	UpdatedBy      int64            `json:"updated_by"`
}

func SetProductCategoryResponse(productCategory models.ProductCategory) ProductCategoryResponse {
	var categories CategoryResponse
	if productCategory.Category != nil {
		categories = SetCategoryResponse(*productCategory.Category)
	}

	return ProductCategoryResponse{
		ID:             productCategory.ID,
		ProductID:      productCategory.ProductID,
		CategoryID:     productCategory.CategoriesID,
		CategoryDetail: categories,
		CreatedAt:      productCategory.CreatedAt,
		CreatedBy:      productCategory.CreatedBy,
		UpdatedAt:      productCategory.UpdatedAt,
		UpdatedBy:      productCategory.UpdatedBy,
	}
}
