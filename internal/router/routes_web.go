package router

import (
	"pleasurelove/internal/controllers"
	"pleasurelove/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutesWeb(api fiber.Router, handler *controllers.AuthController) {
	auth := api.Group("/auth")

	auth.Post("/validate", handler.ValidateCredentials)
	auth.Post("/token", handler.GenerateAccessToken)

	auth.Post("/register")
	auth.Post("/guest")

	auth.Post("/logout", middleware.AuthMiddleware(), handler.Logout)
}

func UserRoutesWeb(api fiber.Router, handler *controllers.UserController) {
	userRoute := api.Group("/user")

	userRoute.Get("/:id", handler.GetProfile)
}
