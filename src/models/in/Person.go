package in

type Persons struct {
	ResponseWrapper
	Results []Person `json:"results"`
}

type Person struct {
	Name      string `json:"name"`
	Homeworld string `json:"homeworld"`
}
