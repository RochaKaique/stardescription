package main

import (
	"bff/src/router"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("/etc/")


	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":8080")
}
