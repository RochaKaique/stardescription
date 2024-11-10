package middleware

import (
	"bff/src/authentication"
	"fmt"
	"os"

	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(ctx *fiber.Ctx) error {
	if err := authentication.ValidateToken(ctx); err != nil {
		fmt.Println(err)
		return err
	}
	return ctx.Next()
}

func HandleLog(ctx *fiber.Ctx) error {
	userId, _ := authentication.ExtractUserID(ctx)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).
	With(
		"User-Agent", authentication.ExtractUserAgent(ctx),
		"userId", userId)

	slog.SetDefault(logger)

	return ctx.Next()
}
