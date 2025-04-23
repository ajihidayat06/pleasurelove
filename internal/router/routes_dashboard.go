package router

import (
	"pleasurelove/internal/constanta"
	"pleasurelove/internal/controllers/dashboard"
	"pleasurelove/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router, authController *dashboard.AuthController) {
	auth := api.Group("/auth")

	auth.Post("/validate", authController.ValidateCredentials) // Endpoint pertama
	auth.Post("/token", authController.GenerateAccessToken)    // Endpoint kedua

	auth.Post("/logout", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), authController.LogoutDashboard)
}

func UserRoutesDashboard(api fiber.Router, handler *dashboard.UserDahboardController) {
	//protected routes
	userDashboard := api.Group("/user")
	userDashboard.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionCreate), handler.CreateUserDashboard)
	userDashboard.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), handler.GetListUser)
	userDashboard.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionRead), handler.GetUserByID)
	userDashboard.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionUpdate), handler.UpdateUserByID)
	userDashboard.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuUserActionDelete), handler.DeleteUserByID)
}

func CategoryRoutesdashboard(api fiber.Router, handler *dashboard.CategoryDashboardController) {
	// Protected routes
	category := api.Group("/category")
	category.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuCategoryActionCreate), handler.CreateCategory)
	category.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuCategoryActionRead), handler.GetListCategory)
	category.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuCategoryActionRead), handler.GetCategoryByID)
	category.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuCategoryActionUpdate), handler.UpdateCategoryByID)
	category.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuCategoryActionDelete), handler.DeleteCategoryByID)
}

func RoleRoutesDashboard(api fiber.Router, handler *dashboard.RoleController) {
	// Protected routes
	role := api.Group("/role")
	role.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuRoleActionCreate), middleware.CheckAdminRoleMiddleware(), handler.CreateRole)
	role.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuRoleActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetListRole)
	role.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRoleActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetRoleByID)
	role.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRoleActionUpdate), middleware.CheckAdminRoleMiddleware(), handler.UpdateRoleByID)
	role.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRoleActionDelete), middleware.CheckAdminRoleMiddleware(), handler.DeleteRoleByID)
}

func PermissionRoutesDashboard(api fiber.Router, handler *dashboard.PermissionController) {
	// Protected routes
	permission := api.Group("/permission")
	permission.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuPermissionsActionCreate), middleware.CheckAdminRoleMiddleware(), handler.CreatePermission)
	permission.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuPermissionsActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetListPermission)
	permission.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuPermissionsActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetPermissionByID)
	permission.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuPermissionsActionUpdate), middleware.CheckAdminRoleMiddleware(), handler.UpdatePermissionByID)
	permission.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuPermissionsActionDelete), middleware.CheckAdminRoleMiddleware(), handler.DeletePermissionByID)
}

func RolePermissionsRoutesDashboard(api fiber.Router, handler *dashboard.RolePermissionsController) {
	// Protected routes
	rolePermissions := api.Group("/role-permissions")
	rolePermissions.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuRolePermissionsActionCreate), middleware.CheckAdminRoleMiddleware(), handler.CreateRolePermission)
	rolePermissions.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuRolePermissionsActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetListRolePermissions)
	rolePermissions.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRolePermissionsActionRead), middleware.CheckAdminRoleMiddleware(), handler.GetRolePermissionByID)
	rolePermissions.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRolePermissionsActionUpdate), middleware.CheckAdminRoleMiddleware(), handler.UpdateRolePermissionByID)
	rolePermissions.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuRolePermissionsActionDelete), middleware.CheckAdminRoleMiddleware(), handler.DeleteRolePermissionByID)
}

func ProductRoutesdashboard(api fiber.Router, handler *dashboard.ProductDashboardController) {
	// Protected routes
	category := api.Group("/product")
	category.Post("/", middleware.AuthMiddlewareDashboard(constanta.MenuProductActionCreate), handler.CreateProduct)
	category.Get("/", middleware.AuthMiddlewareDashboard(constanta.MenuProductActionRead), handler.GetListProduct)
	category.Get("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuProductActionRead), handler.GetProductByID)
	category.Put("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuProductActionUpdate), handler.UpdateProductByID)
	category.Delete("/:id", middleware.AuthMiddlewareDashboard(constanta.MenuProductActionDelete), handler.DeleteProductByID)
}
