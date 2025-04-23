package models

import (
	"time"
)

type ProductCategory struct {
	ID           int64     `gorm:"primaryKey;column:id"`
	ProductID    int64     `gorm:"column:product_id"`
	CategoriesID int64     `gorm:"column:categories_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	CreatedBy    int64     `gorm:"column:created_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	UpdatedBy    int64     `gorm:"column:updated_by"`
	Category     *Category `gorm:"foreignKey:CategoriesID;references:ID"`
}

// TableName overrides the default table name
func (ProductCategory) TableName() string {
	return "product_categories"
}
