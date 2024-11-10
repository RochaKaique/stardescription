package router

import (
	"bff/src/controllers"
	"bff/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/person", middleware.Authenticate, middleware.HandleLog, controllers.GetDescription)
	app.Post("/user", controllers.CreateUser)
	app.Post("/login", middleware.HandleLog, controllers.Login)
}