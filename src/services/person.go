package services

import (
	"bff/src/models/in"
	"errors"

	"github.com/gofiber/fiber/v3/client"
)

func GetPersonByName(name string) (in.Person, error) {
	if name == "" {
		return in.Person{}, errors.New("O nome do personagem não pode estar em branco")
	}
	// resp, error := http.Get("https://swapi.dev/api/people?search="+name)

	cc := client.New()
	resp, error := cc.Get("https://swapi.dev/api/people?search=" + name)
	if error != nil {
		return in.Person{}, error
	}
	defer resp.Close()

	var personData in.PersonData

	if err := resp.JSON(&personData); err != nil {
		return in.Person{}, errors.New("O erro ao ler json")
	}

	return personData.Persons[0], nil
}
