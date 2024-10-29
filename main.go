package main

import (
	database "bff/src/db"
	"bff/src/router"
	"fmt"

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
	db, error := database.Connect()
	if error != nil {
		panic(error)
	} else {
		fmt.Println("conectado com sucess")
	}
	defer db.Close()
	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":8080")
}
