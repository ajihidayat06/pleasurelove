package response

import (
	"pleasurelove/internal/models"
	"time"
)

type UserResponse struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	RoleID    int64          `json:"role_id"`
	RoleName  string         `json:"role_name"`
	BranchID  int64          `json:"branch_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedBy int64          `json:"created_by"`
	UpdatedBy int64          `json:"updated_by"`
	Roles     *RolesResponse `json:"roles,omitempty"`
}

func SetListResponse(user models.User) UserResponse {
	var roleName string
	if user.Roles != nil {
		roleName = user.Roles.Name
	} else {
		roleName = "-" // Atur nilai default jika role tidak di-set
	}
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		RoleID:    user.RoleID,
		RoleName:  roleName,
		BranchID:  user.BranchID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
		Roles:     nil,
	}
}

func SetResponseListUser(users []models.User) []UserResponse {
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, SetListResponse(user))
	}
	return userResponses
}

func SetUserDetailResponse(user models.User) UserResponse {
	var roleName string
	var roles *RolesResponse

	if user.Roles != nil {
		roleName = user.Roles.Name
		r := SetRolesResponse(*user.Roles)
		roles = &r
	} else {
		roleName = "-"
		// Tetap inisialisasi agar "roles" muncul di JSON (walau kosong)
		roles = &RolesResponse{}
	}

	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		RoleID:    user.RoleID,
		RoleName:  roleName,
		BranchID:  user.BranchID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
		Roles:     roles,
	}
}
