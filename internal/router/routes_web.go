package router

import (
	"github.com/gofiber/fiber/v2"
	"pleasurelove/internal/controllers"
)

func UserRoutesWeb(api fiber.Router, handler *controllers.UserController) {
	userRoute := api.Group("/user")

	userRoute.Get("/:id", handler.GetProfile)
}
