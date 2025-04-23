package response

import (
	"pleasurelove/internal/models"
	"time"
)

type CategoryResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int64     `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int64     `json:"updated_by"`
}

func SetCategoryResponse(category models.Category) CategoryResponse {
	return CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Code:      category.Code,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
		CreatedBy: category.CreatedBy,
		UpdatedAt: category.UpdatedAt,
		UpdatedBy: category.UpdatedBy,
	}
}

func SetResponseListCategory(category []models.Category) []CategoryResponse {
	var categoryResponse []CategoryResponse
	for _, cat := range category {
		categoryResponse = append(categoryResponse, SetCategoryResponse(cat))
	}
	return categoryResponse
}
