package authentication

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func CreateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("secret.key")))
}

func ValidateToken(ctx *fiber.Ctx) error {
	tokenString, err := extractToken(ctx)
	if err != nil {
		return err
	}
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func extractToken(ctx *fiber.Ctx) (string, error) {
	token := ctx.GetReqHeaders()["Authorization"]

	if token == nil || len(token) == 0 {
		return "", fiber.NewError(fiber.StatusUnauthorized)
	}

	if len(strings.Split(token[0], " ")) == 2 {
		return strings.Split(token[0], " ")[1], nil
	}

	return "", fiber.NewError(fiber.StatusUnauthorized)
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return []byte(viper.GetString("secret.key")), nil
}
