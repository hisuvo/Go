package main

import "fmt"

func main() {
	fruits := [4]string{"Apple", "Banana", "Mango", "Orange"}

	for i:=0; i < len(fruits); i++{
		fmt.Printf("value is %v\n",fruits[i])
	}

	for _,value := range fruits {
		fmt.Printf("Value is -> %v and this type is -> %T\n", value, value)
	}
}