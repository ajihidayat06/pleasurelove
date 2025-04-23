package request

var ReqPermissionErrorMessage = map[string]string{
	"code":         "Code is required",
	"name":         "Name is required",
	"group_menu":   "GroupMenu is required",
	"action":       "Action is required",
	"access_scope": "AccessScope is required",
}

type ReqPermission struct {
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	GroupMenu   string `json:"group_menu" validate:"required"`
	Action      string `json:"action" validate:"required"`
	AccessScope string `json:"access_scope" validate:"required"`
	AbstractRequest
}

type ReqPermissionUpdate struct {
	ID          int64  `json:"id"`
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	GroupMenu   string `json:"group_menu" validate:"required"`
	Action      string `json:"action" validate:"required"`
	AccessScope string `json:"access_scope" validate:"required"`
	AbstractRequest
}

var ReqPermissionUpdateErrorMessage = map[string]string{
	"Code":          "Code is required",
	"Name":          "Name is required",
	"GroupMenu":     "GroupMenu is required",
	"Action":        "Action is required",
	"AccessScope":   "AccessScope is required",
	"UpdateddAtStr": "UpdatedAt is required",
}
