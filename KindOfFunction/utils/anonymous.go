package utils

import (
	"fmt"
)

func Anonymous() {
	// anonymous function
	func (){
		fmt.Println("This is anonymous function")
	}() // IIFE -> immediately invoked function expression
}
