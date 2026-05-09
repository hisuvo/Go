package utils

import "fmt"

func Expression() {
	// variable expression
	a := 5
	b := 6

	// this is function expression
	add := func(num int, num1 int) {
		sum := num + num1
		fmt.Println("Totla sum is ->",sum)
	}

	add(a+5, b)
}