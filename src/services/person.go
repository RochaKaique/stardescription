package services

import (
	"bff/src/models/in"
	"bff/src/models/out"
	"errors"
	"sync"

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

	chPlanet := make(chan in.Planet)
	chError := make(chan error)

	go getHomeworld(person.Homeworld, chPlanet, chError)
	films, err := fetchFilms(person.Films)
	if err != nil {
		return out.Person{}, err
	}

	var personResponse out.Person

	planet := <-chPlanet
	err = <-chError
	if err != nil {
		return out.Person{}, err
	}

	personResponse.Films = films
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

func getFilm(url string, wg *sync.WaitGroup, chFilm chan<- in.Film, chErr chan<- error) {
	defer wg.Done()
	if url == "" {
		chErr <- errors.New("A uri para busca do filme está vazia")
		return
	}

	cc := client.New()
	resp, err := cc.Get(url)
	if err != nil || resp.StatusCode() != 200 {
		chErr <- errors.New("Erro ao obter filmes")
	}
	defer resp.Close()

	var film in.Film
	if err := resp.JSON(&film); err != nil {
		chErr <- errors.New("Erro ao ler json de resposta")
	}

	chFilm <- film
}

func fetchFilms(urls []string) ([]in.Film, error) {
	var wg sync.WaitGroup
	chFilm := make(chan in.Film, len(urls))
	chErr := make(chan error, 1)

	for _, url := range urls {
		wg.Add(1)
		go getFilm(url, &wg, chFilm, chErr)
	}

	go func() {
		wg.Wait()
		close(chErr)
		close(chFilm)
	}()

	err := <-chErr

	if err != nil {
		return nil, err
	}

	var films []in.Film

	for film := range chFilm {
		films = append(films, film)
	}

	return films, nil
}
