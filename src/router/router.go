package router

import (
	"bff/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/person", controllers.GetDescription)
	app.Post("/user", controllers.CreateUser)
}