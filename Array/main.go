package main

import "fmt"


func main() {
	var numbers = [5]int{}

	persons := [3]string{}

	allPersons := [...] string {"a",3:"b","c","d"}

	fmt.Println(numbers,persons,allPersons)


	var price = [3] int {10,20,30};
	price[1] = 60
	fmt.Println(len(price))
	fmt.Printf("%T",numbers)

}