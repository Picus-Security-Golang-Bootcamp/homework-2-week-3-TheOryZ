package books

import (
	m "Picus-Security-Golang-Bootcamp/homework-2-week-3-TheOryZ/pkg/model"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
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

//SearchWithTitle Function that searches for a book based on the entered value
func SearchWithTitle(s string) (*m.Books, error) {
	var bookList m.Books

	jsonFile, err := os.Open("../../pkg/store/db/books.json")
	if err != nil {
		return nil, err
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var books m.Books
		json.Unmarshal(byteValue, &books)
		for i := 0; i < len(books.Books); i++ {
			if strings.Contains(strings.ToLower(books.Books[i].Title), strings.ToLower(s)) {
				bookList.Books = append(bookList.Books, books.Books[i])
			}
		}
	}
	defer jsonFile.Close()
	return &bookList, nil
}

//DeleteBook Deletes the book based on the sent book id
func DeleteBook(id int) error {
	jsonFile, err := os.Open("../../pkg/store/db/books.json")
	if err != nil {
		return err
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var books m.Books
		json.Unmarshal(byteValue, &books)
		for i := 0; i < len(books.Books); i++ {
			if books.Books[i].ID == id {
				books.Books[i].IsDeleted = true
				break
			}
		}
		byteValue, err = json.Marshal(books)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile("../../pkg/store/db/books.json", byteValue, 0644)
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		return nil
	}
}

//BuyBook Makes the buying with the entered book id and quantity information.
func BuyBook(id, quantity int) error {
	jsonFile, err := os.Open("../../pkg/store/db/books.json")
	if err != nil {
		return err
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var books m.Books
		json.Unmarshal(byteValue, &books)
		for i := 0; i < len(books.Books); i++ {
			if books.Books[i].ID == id {
				books.Books[i].NumberOfStocks = books.Books[i].NumberOfStocks - quantity
				break
			}
		}
		byteValue, err = json.Marshal(books)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile("../../pkg/store/db/books.json", byteValue, 0644)
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		return nil
	}
}

//AvailabilityOfBuying It returns the status of being sold by comparing the sent book id and stock.
func AvailabilityOfBuying(id, quantity int) bool {
	jsonFile, err := os.Open("../../pkg/store/db/books.json")
	if err != nil {
		return false
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var books m.Books
		json.Unmarshal(byteValue, &books)
		for i := 0; i < len(books.Books); i++ {
			if books.Books[i].ID == id {
				if books.Books[i].IsDeleted {
					return false
				}
				if books.Books[i].NumberOfStocks < quantity {
					return false
				}
			}
		}
	}
	defer jsonFile.Close()
	return true
}
