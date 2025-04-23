package router

import (
	"pleasurelove/internal/controllers"
	"pleasurelove/internal/middleware"
	"pleasurelove/pkg/redis"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/health", controllers.HealthCheck(controllers.HealthDependencies{
		DB:    db,
		Redis: redis.RDB,
	}))

	DashboardRoute(app, db)

	WebRoute(app, db)
}

func DashboardRoute(app *fiber.App, db *gorm.DB) {
	auth := InitAuth(db)
	user := InitUserDahboard(db)
	category := InitCategoryDashboard(db)
	role := InitRoleDashboard(db)
	permissions := InitPermissionDashboard(db)
	rolePermissions := InitRolePermissionsDashboard(db)
	product := InitProductDashboard(db)

	api := app.Group("/api/v1/dashboard")
	// Public routes

	AuthRoutes(api, auth)

	RoleRoutesDashboard(api, role)
	PermissionRoutesDashboard(api, permissions)
	RolePermissionsRoutesDashboard(api, rolePermissions)

	UserRoutesDashboard(api, user)
	CategoryRoutesdashboard(api, category)
	ProductRoutesdashboard(api, product)
}

func WebRoute(app *fiber.App, db *gorm.DB) {
	user := InitUser(db)

	api := app.Group("/api/v1")

	api.Post("/login", user.Login)
	api.Post("/register", user.Register)
	api.Post("/logout", middleware.AuthMiddleware(), user.Logout)

	// Protected routes
	UserRoutesWeb(api, user)
}
