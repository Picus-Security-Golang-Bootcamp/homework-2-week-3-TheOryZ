package model

//Authors authors model
type Authors struct {
	Books []Author `json:"Author"`
}

//Author author detail model
type Author struct {
	ID   int    `json:"Id"`
	Name string `json:"Name"`
}
