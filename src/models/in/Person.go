package in

type PersonData struct {
	ResponseWrapper
	Persons []Person `json:"results"`
}

type Person struct {
	Name      string `json:"name"`
	Homeworld string `json:"homeworld"`
}
