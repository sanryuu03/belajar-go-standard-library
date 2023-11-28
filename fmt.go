package main

import "fmt"

func main() {
	firstName := "san"
	lastName := "ryuu"

	fmt.Println("hello '", firstName, lastName, "'")
	fmt.Printf("hello '%s %s'\n", firstName, lastName)
}
