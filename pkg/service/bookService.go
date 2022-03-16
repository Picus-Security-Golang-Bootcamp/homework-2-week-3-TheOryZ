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

//SearchWithTitle Function that searches for a book based on the entered value
func SearchWithTitle(s string) (*m.Book, error) {
	return b.SearchWithTitle(s)
}
