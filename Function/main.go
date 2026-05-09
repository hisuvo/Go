package main

import "fmt"

func show(value any){
	fmt.Println(value)
}

func getNumber(num1 int, num2 int) (int, int){
	add := num1 + num2
	mult := num1 * num2

	return  add, mult
}

func testFunc(name string, age int) (isName string, isAge int) {
	isName = name
	isAge = age
	return
}

func sum(num1 int, num2 int) int {
	total := num1 + num2
	return total
}

func main() {
	a , _ := testFunc("Golang",2007)
	show(a)
	show(sum(4, 6))

	x := 4
	y := 5
	num1, num2 := getNumber(x, y)
	show(num1)
	show(num2)

}