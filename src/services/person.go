package services

import (
	"bff/src/models/in"
	"bff/src/models/out"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v3/client"
)

func GetPersonByName(name string) (out.Person, error) {
	if name == "" {
		return out.Person{}, errors.New("O nome do personagem não pode estar em branco")
	}
	// resp, error := http.Get("https://swapi.dev/api/people?search="+name)

	cc := client.New()
	resp, error := cc.Get("https://swapi.dev/api/people?search=" + name)
	if error != nil {
		return out.Person{}, error
	}
	defer resp.Close()

	var personData in.PersonData

	if err := resp.JSON(&personData); err != nil {
		return out.Person{}, errors.New("O erro ao ler json")
	}

	person := personData.Persons[0]
	planet, err := getHomeworld(person.Homeworld)
	if err != nil {
		return out.Person{}, err
	}

	var personResponse out.Person
	
	personResponse.Name = person.Name
	personResponse.Homeworld = planet.Name

	for _, filmUri := range person.Films {
		film, err := getFilm(filmUri)
		if err != nil {
			return out.Person{}, err
		}
		personResponse.Films = append(personResponse.Films, film)
	}
	
	return personResponse, nil
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

	fmt.Println(string(resp.Body()))

	var planet in.Planet
	if err := resp.JSON(&planet); err != nil {
		return in.Planet{}, errors.New("Erro ao ler json de resposta")
	}

	return planet, nil
}

func getFilm(uri string) (in.Film, error) {
	if uri == "" {
		return in.Film{}, errors.New("A uri para busca da terra natal está vazia")
	}

	cc := client.New()
	resp, err := cc.Get(uri)
	if err != nil {
		return in.Film{}, err
	}
	defer resp.Close()

	fmt.Println(string(resp.Body()))

	var film in.Film
	if err := resp.JSON(&film); err != nil {
		return in.Film{}, errors.New("Erro ao ler json de resposta")
	}

	return film, nil
}
