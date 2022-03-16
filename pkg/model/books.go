package model

//Books books model
type Books struct {
	Books []Book `json:"Books"`
}

//Book book detail model
type Book struct {
	ID             int     `json:"Id"`
	Title          string  `json:"Title"`
	NumberOfPages  int     `json:"Number_Of_Pages"`
	NumberOfStocks int     `json:"Number_Of_Stocks"`
	Price          float64 `json:"Price"`
	ISBN           string  `json:"ISBN"`
	ReleaseDate    string  `json:"Release_Date"`
	Author         Author  `json:"Author"`
	IsDeleted      bool    `json:"IsDeleted"`
}

//Author author detail model
type Author struct {
	ID   int    `json:"Id"`
	Name string `json:"Name"`
}
