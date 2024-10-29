package in

type PersonData struct {
	Count   int      `json:"count"`
	Next    string   `json:"next"`
	Previus string   `json:"previous"`
	Persons []Person `json:"results"`
}

type Person struct {
	Name      string `json:"name"`
	Homeworld string `json:"homeworld"`
	Films     []string `json:"films"`
}
