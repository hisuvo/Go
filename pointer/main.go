package main

import "fmt"

// Note:
// 1. pase by value
// use value copy and pest

// 2. pase by reference
// -> & ambersend or address of value
// -> * value at address


func print(numbers [] int){ // pase infinite value
	fmt.Println(numbers)
}

func printStr(fruits *[3] string){ // pase onlu three value
	fmt.Println(fruits)
}

type Student struct {
	name string
	age int
	salary float64
}

func main() {
	// pointer or address of memory (ram)
	number1 := 20

	// ambersand & => address of number1 | 824,633,825,000
	address := &number1 
	
	// value assign at the number1 address
	*address = 50
	
	fmt.Println("Address of number1", address)
	// * => value at address
	fmt.Println("Value the address of number1 ->", *address)

	arr := [] int{2,3,4,34, 45, 56, 60} // pase infinite value
	print(arr)
	
	arr1 := [3] string {"apple", "mango","Orrange"} // pase only 3 value
	printStr(&arr1)
	fmt.Println(&arr1)

	student1 := Student{
		name: "Suvo Datta",
		age: 24,
		salary: 233456,
	}

	stu := &student1

	fmt.Println("student ->",*stu)
}