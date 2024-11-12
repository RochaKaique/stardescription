package services

import (
	"bff/src/models/in"
	"bff/src/models/out"
	"errors"
	"fmt"
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

	//Chamadas de planeta e filmes
	// planet, err := getHomeworld(person.Homeworld)
	// if err != nil {
	// 	return out.Person{}, err
	// }

	chPlanet := make(chan in.Planet)
	chError := make(chan error)

	go getHomeworld(person.Homeworld, chPlanet, chError)
	films, err := getFilmsAsync(person.Films)
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

func getFilmsAsync(uris []string) ([]in.Film, error) {
	filmCh := make(chan in.Film)
	errorCh := make(chan error)
	done := make(chan struct{})

	uriCh := streamUris(done, uris)

	var wg sync.WaitGroup
	wg.Add(len(uris))

	for i := 0; i < len(uris); i++ {
		go func() {
			for uri := range uriCh {
				film, error := getFilm(uri)
				if error != nil {
					errorCh <- error
				}
				filmCh <- film
				wg.Done()
			}
		}()
	}

	go func() {
        wg.Wait()
        close(filmCh)
		close(errorCh)
    }()

	var films []in.Film

	for film := range filmCh {
		fmt.Println(film)
		films = append(films, film)
	}

	return films, nil
}

func streamUris(done <-chan struct{}, uris []string) <-chan string {
	uriCh := make(chan string)
	go func() {
		defer close(uriCh)
		for _, uri := range uris {
			select {
			case uriCh <- uri:
			case <-done:
				break
			}
		}
	}()
	return uriCh
}
