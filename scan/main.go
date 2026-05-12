package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Welcome ", firstName, lastName)
}