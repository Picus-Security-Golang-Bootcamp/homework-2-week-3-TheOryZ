package main

import (
	s "Picus-Security-Golang-Bootcamp/homework-2-week-3-TheOryZ/pkg/service"
	"fmt"
	"os"
	"strings"
)

func main() {
	s.SetFlags()
	args := os.Args[1:]
	validStatus := s.ValidationCommand(args)
	if !validStatus {
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
		}
	}
}
