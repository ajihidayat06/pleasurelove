package request

type ReqRolePermission struct {
	ID           int64  `json:"id"`
	PermissionID int64  `json:"permission_id" validate:"gt=0"`
	Scope        string `json:"scope"`
	AbstractRequest
}

var ReqRolePermissionErrorMessage = map[string]string{
	"permission_id": "Permission ID required",
}
