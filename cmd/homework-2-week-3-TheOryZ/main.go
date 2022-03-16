package main

import (
	b "Picus-Security-Golang-Bootcamp/homework-2-week-3-TheOryZ/pkg/service"
	"fmt"
	"os"
)

func main() {
	b.SetFlags()
	args := os.Args[1:]
	validationStatus := b.ValidationCommand(args)
	fmt.Println(validationStatus)
	//if validationStatus true do it operations. And then false return alert
}
