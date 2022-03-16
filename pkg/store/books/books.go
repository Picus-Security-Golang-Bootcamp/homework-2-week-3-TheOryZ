package books

import (
	m "Picus-Security-Golang-Bootcamp/homework-2-week-3-TheOryZ/pkg/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

// GetAllBooks Function that returns all books
func GetAllBooks() (*m.Books, error) {
	var books m.Books
	jsonFile, err := os.Open("../../pkg/store/db/books.json")
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &books)
	defer jsonFile.Close()
	return &books, nil
}

//GetById Function that returns book by entered id
func GetById(id int) (*m.Book, error) {
	var book m.Book
	jsonFile, err := os.Open("../../pkg/store/db/books.json")
	if err != nil {
		return nil, err
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var books m.Books
		json.Unmarshal(byteValue, &books)
		for i := 0; i < len(books.Books); i++ {
			if books.Books[i].ID == id {
				book.ID = books.Books[i].ID
				book.Title = books.Books[i].Title
				book.NumberOfPages = books.Books[i].NumberOfPages
				book.NumberOfStocks = books.Books[i].NumberOfStocks
				book.Price = books.Books[i].Price
				book.ISBN = books.Books[i].ISBN
				book.ReleaseDate = books.Books[i].ReleaseDate
				book.Author = books.Books[i].Author
				book.IsDeleted = books.Books[i].IsDeleted
				break
			}
		}
	}
	defer jsonFile.Close()
	return &book, nil
}
