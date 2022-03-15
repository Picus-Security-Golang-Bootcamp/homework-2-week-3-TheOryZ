package service

import (
	model "Picus-Security-Golang-Bootcamp/homework-2-week-3-TheOryZ/pkg/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// GetAllBooks Function that returns all books
func GetAllBooks() model.Books {
	var books model.Books
	jsonFile, err := os.Open("pkg/db/books.json")
	if err != nil {
		fmt.Println(err)
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &books)
	}
	defer jsonFile.Close()
	return books
}

//GetById Function that returns book by entered id
func GetById(id int) (model.Book, error) {
	var book model.Book
	jsonFile, err := os.Open("pkg/db/books.json")
	if err != nil {
		return book, err
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var books model.Books
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
				book.AuthorID = books.Books[i].AuthorID
				book.IsDeleted = books.Books[i].IsDeleted
				break
			}
		}
	}
	defer jsonFile.Close()
	return book, nil
}
