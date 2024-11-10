package controllers

import (
	"bff/src/services"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func GetDescription(ctx *fiber.Ctx) error {
	queries := ctx.Queries()
	name := queries["name"]

	person, _ := services.GetPersonByName(name)

	slog.Info("Consulta Realizada")

	return ctx.JSON(person)
}
