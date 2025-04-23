package models

import "time"

type Roles struct {
	ID              int64              `json:"id"`
	Code            string             `json:"code"`
	Name            string             `json:"name"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	CreatedBy       int64              `json:"created_by"`
	UpdatedBy       int64              `json:"updated_by"`
	RolePermissions *[]RolePermissions `json:"role_permissions" gorm:"foreignKey:RoleID"`
}

func (Roles) TableName() string {
	return "roles"
}