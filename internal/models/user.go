package models

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	RoleID    int64     `json:"role_id"`
	BranchID  int64     `json:"branch_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy int64     `json:"created_by"`
	UpdatedBy int64     `json:"updated_by"`
	Roles     *Roles    `json:"roles" gorm:"foreignKey:RoleID"`
}

func (User) Tablename() string {
	return "users"
}

type UserLogin struct {
	ID              int64             `json:"id"`
	RoleID          int64             `json:"role_id"`
	RoleName        string            `json:"role_name"`
	RoleCode        string            `json:"role_code"`
	RolePermissions []RolePermissions `json:"permissions"` // Gunakan RolePermissions di sini
}

func (UserLogin) TableName() string {
	return "users"
}
