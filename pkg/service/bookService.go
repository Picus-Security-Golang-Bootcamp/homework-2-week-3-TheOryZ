package service

import (
	m "Picus-Security-Golang-Bootcamp/homework-2-week-3-TheOryZ/pkg/model"
	b "Picus-Security-Golang-Bootcamp/homework-2-week-3-TheOryZ/pkg/store/books"
)

// GetAllBooks Function that returns all books
func GetAllBooks() (*m.Books, error) {
	return b.GetAllBooks()
}

//GetById Function that returns book by entered id
func GetById(id int) (*m.Book, error) {
	return b.GetById(id)
}

//SearchWithTitle Function that searches for a book based on the entered value (Without Deleted items)
func SearchWithTitle(s string) (*m.Books, error) {
	return b.SearchWithTitle(s)
}

//DeleteBook
func DeleteBook(id int) error {
	return b.DeleteBook(id)
}

//BuyBook Makes the buying with the entered book id and quantity information.
func BuyBook(id, quantity int) error {
	return b.BuyBook(id, quantity)
}

//AvailabilityOfBuying It returns the status of being sold by comparing the sent book id and stock.
func AvailabilityOfBuying(id, quantity int) bool {
	return b.AvailabilityOfBuying(id, quantity)
}
