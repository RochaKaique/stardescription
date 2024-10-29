package out

import "bff/src/models/in"

type Person struct {
	Name      string `json:"name"`
	Homeworld string `json:"homeworld"`
	Films     []in.Film `json:"films"`
}