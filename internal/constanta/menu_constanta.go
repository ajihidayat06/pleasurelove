package constanta

const (
	MenuGroupUser            = "user"
	MenuGroupCategory        = "category"
	MenuGroupRole            = "role"
	MenuGroupPermissions     = "permissions"
	MenuGroupRolePermissions = "role_permissions"
	MenuGroupProduct         = "product"
)

const (
	MenuUserActionCreate = MenuGroupUser + ":" + AuthActionCreate
	MenuUserActionRead   = MenuGroupUser + ":" + AuthActionRead
	MenuUserActionUpdate = MenuGroupUser + ":" + AuthActionUpdate
	MenuUserActionDelete = MenuGroupUser + ":" + AuthActionDelete

	MenuCategoryActionCreate = MenuGroupCategory + ":" + AuthActionCreate
	MenuCategoryActionRead   = MenuGroupCategory + ":" + AuthActionRead
	MenuCategoryActionUpdate = MenuGroupCategory + ":" + AuthActionUpdate
	MenuCategoryActionDelete = MenuGroupCategory + ":" + AuthActionDelete

	MenuRoleActionCreate = MenuGroupRole + ":" + AuthActionCreate
	MenuRoleActionRead   = MenuGroupRole + ":" + AuthActionRead
	MenuRoleActionUpdate = MenuGroupRole + ":" + AuthActionUpdate
	MenuRoleActionDelete = MenuGroupRole + ":" + AuthActionDelete

	MenuPermissionsActionCreate = MenuGroupPermissions + ":" + AuthActionCreate
	MenuPermissionsActionRead   = MenuGroupPermissions + ":" + AuthActionRead
	MenuPermissionsActionUpdate = MenuGroupPermissions + ":" + AuthActionUpdate
	MenuPermissionsActionDelete = MenuGroupPermissions + ":" + AuthActionDelete

	MenuRolePermissionsActionRead   = MenuGroupRolePermissions + ":" + AuthActionRead
	MenuRolePermissionsActionCreate = MenuGroupRolePermissions + ":" + AuthActionCreate
	MenuRolePermissionsActionUpdate = MenuGroupRolePermissions + ":" + AuthActionUpdate
	MenuRolePermissionsActionDelete = MenuGroupRolePermissions + ":" + AuthActionDelete

	MenuProductActionCreate = MenuGroupProduct + ":" + AuthActionCreate
	MenuProductActionRead   = MenuGroupProduct + ":" + AuthActionRead
	MenuProductActionUpdate = MenuGroupProduct + ":" + AuthActionUpdate
	MenuProductActionDelete = MenuGroupProduct + ":" + AuthActionDelete
)

const (
	FieldUserID      = "ID"
	FieldCode        = "CODE"
	FieldName        = "NAME"
	FieldPermissions = "PERMISSIONS"
	FieldCategory    = "CATEGORY"
)
