package in

type ResponseWrapper struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Previus string `json:"previous"`
}
