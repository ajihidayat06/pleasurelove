package request

type ReqRoles struct {
	Code            string              `json:"code"`
	Name            string              `json:"name"`
	RolePermissions []ReqRolePermission `json:"role_permissions" validate:"dive"`
}

var ReqRolesErrorMessage = map[string]string{
	"Code":         "code required",
	"Name":         "name required",
	"PermissionID": "Permission ID required",
}

type ReqRoleUpdate struct {
	ID              int64               `json:"id" validate:"required"`
	Code            string              `json:"code" validate:"required"`
	Name            string              `json:"name" validate:"required"`
	RolePermissions []ReqRolePermission `json:"role_permissions" validate:"required,dive"`
	AbstractRequest
}

var ReqRoleUpdateErrorMessage = map[string]string{
	"ID":           "id required",
	"Code":         "code required",
	"Name":         "name required",
	"UpdateddAtStr":    "updated_at required",
	"PermissionID": "Permission ID required",
}
