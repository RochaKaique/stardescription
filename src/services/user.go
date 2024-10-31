package services

import (
	database "bff/src/db"
	"bff/src/models/in"
	"bff/src/repository"
)

func CreateUser(user in.User) error {

	if err := user.Prepare("cadastro"); err != nil {
		return err
	}

	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	_, err = userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}
