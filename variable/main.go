package main

import "fmt"


const A int = 345

func main(){
	// Key type is specified
	var num1, num2, num3, num4 int = 2, 5, 3, 5
	fmt.Println(num1,num2,num3,num4)

	// If key type is not specified
	var name1 , age1 = "Rajib",18
	name2, age2 := "Sajeeb", 24
	fmt.Println(name1, age1)
	fmt.Println(name2, age2)

	// Variable declaration in a block
	var(
		age int = 24
		name string = "Suvo datta"
		country string = "Bangladesh"
		isMarried bool 
	)

	fmt.Println(age, name, country, isMarried)
	fmt.Println(A)
}
