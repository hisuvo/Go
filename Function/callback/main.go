package main

import "fmt"

func middleware(num1 int, num2 int, callback func()) func() {
	return func() {
		fmt.Println("number one", num1)
		callback()
		fmt.Println("number two", num2)
	}
}

func callback() {
		fmt.Println("Main functionn is running")
	}

func main() {
	newResult := middleware(2, 3, callback)
	newResult()
}