package main

import "fmt"

func change(name *string) {
	// fmt.Println("My name is", *name)
	*name = "Suvo"
}

func main() {
	myName := "Rahim"
	yourName := "Karin"
	change(&myName)
	change(&yourName)

	fmt.Println("myname",myName)
	fmt.Println("yourname",yourName)
}
