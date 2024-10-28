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

	person := personData.Persons[0]
	planet, err := getHomeworld(person.Homeworld)
	if err != nil {
		return in.Person{}, err
	}
	person.Homeworld = planet.Name

	return person, nil
}

func getHomeworld(uri string) (in.Planet, error) {
	if uri == "" {
		return in.Planet{}, errors.New("A uri para busca da terra natal está vazia")
	}

	cc := client.New()
	resp, err := cc.Get(uri)
	if err != nil {
		return in.Planet{}, err
	}
	defer resp.Close()

	var planet in.Planet
	if err := resp.JSON(&planet); err != nil {
		return in.Planet{}, errors.New("Erro ao ler json de resposta")
	}

	return planet, nil
}
