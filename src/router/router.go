package router

import (
	"bff/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/hello", controllers.Hello)
}