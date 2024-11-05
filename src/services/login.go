package services

import (
	"bff/src/authentication"
	database "bff/src/db"
	"bff/src/models/in"
	"bff/src/models/out"
	"bff/src/repository"
	"bff/src/security"

	"github.com/gofiber/fiber/v3"
)

func Login(user *in.User) (out.Login, error) {
	db, err := database.Connect()
	if err != nil {
		return out.Login{}, fiber.NewError(fiber.StatusInternalServerError, "Erro ao conectar com obanco: " + err.Error()) 
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userFromDb, err := userRepo.FindByEmailForLogin(user.Email)
	if err != nil {
		return out.Login{}, fiber.NewError(fiber.StatusInternalServerError, "Erro ao obter usuário: " + err.Error()) 
	}

	if err = security.CheckPassword(userFromDb.Password, user.Password); err != nil {
		return out.Login{}, fiber.NewError(fiber.StatusForbidden, "Senha atual não corresponde a cadastrada no banco")
	}

	token, err := authentication.CreateToken(userFromDb.ID)
	if err != nil {
		return out.Login{}, fiber.NewError(fiber.StatusForbidden, "Erro ao criar token: " + err.Error())
	}

	login := new(out.Login)
	login.AccessToken = token

	return *login, nil
}
