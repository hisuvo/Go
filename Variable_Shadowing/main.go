package main

import "fmt"

var a = 10

func main() {
	age := 30
	
	if age > 120 {
		a := 47
		fmt.Println("Condition is ture ->",a)
	} else {
		a := 40
		fmt.Println("Condition is false ->",a)
	}

	fmt.Println("Contition outSide ->",a)
}