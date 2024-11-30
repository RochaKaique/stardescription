package controllers

import (
	"bff/src/services"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func GetDescription(ctx *fiber.Ctx) error {
	queries := ctx.Queries()
	name := queries["name"]

	person, err := services.GetPersonByName(name)
	if err != nil {
		slog.Error("Problema ao realizar consulta")
		return err
	}

	slog.Info("Consulta Realizada")

	return ctx.JSON(person)
}
