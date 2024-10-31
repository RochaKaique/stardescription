package controllers

import (
	"bff/src/models/in"
	"bff/src/services"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	u := new(in.User)
	if err := ctx.BodyParser(u); err != nil {
		return err
	}

	if err := services.CreateUser(*u); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}
