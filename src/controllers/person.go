package controllers

import (
	"bff/src/services"

	"github.com/gofiber/fiber/v2"
)

func GetDescription(ctx *fiber.Ctx) error {
	queries := ctx.Queries()
	name := queries["name"]

	person, _ := services.GetPersonByName(name)

	return ctx.JSON(person)
}
