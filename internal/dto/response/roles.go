package response

import (
	"pleasurelove/internal/models"
	"time"
)

type RolesResponse struct {
	ID              int64                     `json:"id"`
	Code            string                    `json:"code"`
	Name            string                    `json:"name"`
	CreatedAt       time.Time                 `json:"created_at"`
	CreatedBy       int64                     `json:"created_by"`
	UpdatedAt       time.Time                 `json:"updated_at"`
	UpdatedBy       int64                     `json:"updated_by"`
	RolePermissions []RolePermissionsResponse `json:"role_permissions"`
}

func SetRolesResponse(user models.Roles) RolesResponse {
	return RolesResponse{
		ID:        user.ID,
		Name:      user.Name,
		Code:      user.Code,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
	}
}

func SetListResponseRole(user []models.Roles) []RolesResponse {
	var roleResponses []RolesResponse
	for _, role := range user {
		roleResponses = append(roleResponses, SetRolesResponse(role))
	}
	return roleResponses
}

func SetRoleDetailResponse(role models.Roles) RolesResponse {
	var rolePermissions []RolePermissionsResponse
	for _, rolePermission := range *role.RolePermissions {
		var permissions *PermissionsResponse
		if rolePermission.Permissions == nil {
			permissions = nil
		} else {
			permissions = SetPermissionsRespons(*rolePermission.Permissions)
		}

		rp := RolePermissionsResponse{
			ID:           rolePermission.ID,
			RoleID:       rolePermission.RoleID,
			PermissionID: rolePermission.PermissionsID,
			Scope:        rolePermission.AccessScope,
			CreatedAt:    rolePermission.CreatedAt,
			UpdatedAt:    rolePermission.UpdatedAt,
			CreatedBy:    rolePermission.CreatedBy,
			UpdatedBy:    rolePermission.UpdatedBy,
			Permissions:  permissions,
		}
		rolePermissions = append(rolePermissions, rp)
	}
	return RolesResponse{
		ID:              role.ID,
		Name:            role.Name,
		Code:            role.Code,
		CreatedAt:       role.CreatedAt,
		UpdatedAt:       role.UpdatedAt,
		CreatedBy:       role.CreatedBy,
		UpdatedBy:       role.UpdatedBy,
		RolePermissions: rolePermissions,
	}
}
