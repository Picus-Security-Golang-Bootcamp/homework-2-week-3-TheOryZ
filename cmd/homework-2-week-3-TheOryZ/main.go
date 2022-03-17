package main

import (
	s "Picus-Security-Golang-Bootcamp/homework-2-week-3-TheOryZ/pkg/service"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s.SetFlags()
	args := os.Args[1:]
	validStatus := s.ValidationCommand(args)
	if validStatus != nil {
		fmt.Println("Please enter a valid command. You can see the commands by typing -help.")
	} else {
		command := strings.ToLower(args[0])
		if command == "list" {
			books, err := s.GetAllBooks()
			if err != nil {
				fmt.Println("An error occurred while executing the GetAllBooks function. Error :", err.Error())
			} else {
				for i := 0; i < len(books.Books); i++ {
					fmt.Println("Book Title: " + books.Books[i].Title)
					fmt.Printf("Book Price: %v\n", books.Books[i].Price)
					fmt.Printf("Book Stocks: %v\n", books.Books[i].NumberOfStocks)
					fmt.Println(strings.Repeat("*", 25))
				}
			}
		} else if command == "search" {
			bookTitle := args[1:]
			search := strings.Join(bookTitle, " ")
			book, err := s.SearchWithTitle(search)
			if err != nil {
				fmt.Println("An error occurred while executing the SearchWithTitle function. Error :", err.Error())
			} else if len(book.Title) <= 0 {
				fmt.Println(strings.Repeat("*", 25))
				fmt.Println("No information was found for the book you were looking for.")
				fmt.Println(strings.Repeat("*", 25))
			} else {
				fmt.Println("Book Title: " + book.Title)
				fmt.Printf("Book Price: %v\n", book.Price)
				fmt.Printf("Book Stocks: %v\n", book.NumberOfStocks)
				fmt.Println(strings.Repeat("*", 25))
			}
		} else if command == "get" {
			bookId, _ := strconv.Atoi(args[1])
			book, err := s.GetById(bookId)
			if err != nil {
				fmt.Println("An error occurred while executing the GetById function. Error :", err.Error())
			} else if len(book.Title) <= 0 {
				fmt.Println(strings.Repeat("*", 25))
				fmt.Println("No information was found for the book you were looking for.")
				fmt.Println(strings.Repeat("*", 25))
			} else {
				fmt.Println("Book Title: " + book.Title)
				fmt.Printf("Book Price: %v\n", book.Price)
				fmt.Printf("Book Stocks: %v\n", book.NumberOfStocks)
				fmt.Printf("Book is Deleted: %v\n", book.IsDeleted)
				fmt.Println(strings.Repeat("*", 25))
			}
		} else if command == "delete" {
			bookId, _ := strconv.Atoi(args[1])
			err := s.DeleteBook(bookId)
			if err != nil {
				fmt.Println("An error occurred while executing the DeleteBook function. Error :", err.Error())
			} else {
				fmt.Println("Book deletion successful")
				fmt.Println(strings.Repeat("*", 25))
			}
		}
	}
}
