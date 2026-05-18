package main

import "fmt"

func print[T any](value T) {
	fmt.Println(value)
}

func main() {
	print("Hello 'SUVO DATTA'")
}