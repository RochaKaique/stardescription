package router

import (
	"bff/src/controllers"
	"bff/src/middleware"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	prometheus := fiberprometheus.New("stardescription")
	prometheus.RegisterAt(app, "/metrics")
	prometheus.SetSkipPaths([]string{"/ping"})
	
	app.Use(prometheus.Middleware)
	app.Use(healthcheck.New())

	app.Get("/person", middleware.Authenticate, middleware.HandleLog, controllers.GetDescription)
	app.Post("/user", controllers.CreateUser)
	app.Post("/login", middleware.HandleLog, controllers.Login)
}