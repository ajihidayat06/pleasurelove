package models

import "time"

type Permissions struct {
	ID          int64     `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	GroupMenu   string    `json:"group_menu"`
	Action      string    `json:"action"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   int64     `json:"created_by"`
	UpdatedBy   int64     `json:"updated_by"`
}

func (Permissions) TableName() string {
	return "permissions"
}
type RolePermissions struct {
	ID            int64       `json:"id"`
	RoleID        int64       `json:"role_id"`
	PermissionsID int64       `json:"permissions_id"`
	AccessScope   string      `json:"access_scope"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	CreatedBy     int64       `json:"created_by"`
	UpdatedBy     int64       `json:"updated_by"`
	Permissions   *Permissions `json:"permissions" gorm:"foreignKey:PermissionsID"`
}

func (RolePermissions) TableName() string {	
	return "role_permissions"
}