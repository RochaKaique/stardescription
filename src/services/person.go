package services

import (
	"bff/src/models/in"
	"bff/src/models/out"
	"errors"

	"github.com/gofiber/fiber/v3/client"
)

func GetPersonByName(name string) (out.Person, error) {
	if name == "" {
		return out.Person{}, errors.New("O nome do personagem não pode estar em branco")
	}
	// resp, error := http.Get("https://swapi.dev/api/people?search="+name)

	cc := client.New()
	resp, err := cc.Get("https://swapi.dev/api/people?search=" + name)
	if err != nil {
		return out.Person{}, err
	}
	defer resp.Close()

	var personData in.PersonData

	if err := resp.JSON(&personData); err != nil {
		return out.Person{}, errors.New("O erro ao ler json")
	}

	person := personData.Persons[0]

	//Chamadas de planeta e filmes
	// planet, err := getHomeworld(person.Homeworld)
	// if err != nil {
	// 	return out.Person{}, err
	// }

	chPlanet := make(chan in.Planet)
	chError := make(chan error)

	go getHomeworld(person.Homeworld, chPlanet, chError)

	var personResponse out.Person

	for _, filmUri := range person.Films {
		film, err := getFilm(filmUri)
		if err != nil {
			return out.Person{}, err
		}
		personResponse.Films = append(personResponse.Films, film)
	}

	planet := <-chPlanet
	err = <-chError
	if err != nil {
		return out.Person{}, err
	}

	personResponse.Name = person.Name
	personResponse.Homeworld = planet.Name

	return personResponse, nil
}

func getHomeworld(uri string, chPlanet chan in.Planet, chError chan error) {
	if uri == "" {
		chPlanet <- in.Planet{}
		chError <- errors.New("A uri para busca da terra natal está vazia")
	}

	cc := client.New()
	resp, err := cc.Get(uri)
	if err != nil {
		chPlanet <- in.Planet{}
		chError <- err
	}
	defer resp.Close()

	var planet in.Planet
	if err := resp.JSON(&planet); err != nil {
		chPlanet <- in.Planet{}
		chError <- errors.New("Erro ao ler json de resposta")
	}

	chPlanet <- planet
	chError <- nil
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


	var film in.Film
	if err := resp.JSON(&film); err != nil {
		return in.Film{}, errors.New("Erro ao ler json de resposta")
	}

	return film, nil
}
