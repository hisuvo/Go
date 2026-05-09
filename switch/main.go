//? Expression
//* An expression is something that produces a value.

package main

import "fmt"

func main() {
	switch num := 1; num%2{
	case  0:
		fmt.Println("This is even number")

	default:
		fmt.Println("This is odd number")
	}
}