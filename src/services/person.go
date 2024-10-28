package services

import (
	"bff/src/models/in"
	"errors"

	"github.com/gofiber/fiber/v3/client"
)

func GetPersonByName(name string) (in.Persons, error) {
	if name == "" {
		return in.Persons{}, errors.New("O nome do personagem não pode estar em branco")
	}
	// resp, error := http.Get("https://swapi.dev/api/people?search="+name)

	cc := client.New()
	resp, error := cc.Get("https://swapi.dev/api/people?search=" + name)
	if error != nil {
		return in.Persons{}, error
	}
	defer resp.Close()

	var persons in.Persons

	if err := resp.JSON(&persons); err != nil {
		return in.Persons{}, errors.New("O erro ao ler json")
	}

	return persons, nil
}
