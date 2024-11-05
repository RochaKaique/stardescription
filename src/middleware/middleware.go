package middleware

import (
	"bff/src/authentication"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(ctx *fiber.Ctx) error {
	if err := authentication.ValidateToken(ctx); err != nil {
		fmt.Println(err)
		return err
	}
	return ctx.Next()
}