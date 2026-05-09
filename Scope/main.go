package main

import "fmt"

var (
	a int = 10
	b int = 20
)

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func mult(num1 int, num2 int) int {
	sum := num1 * num2
	return sum
}

func getAbout(name string, age int, job string)  {
	 fmt.Println(name,"is", age,"Years old and He is a", job)
}

func main() {
	x := 34
	y := 16

	z := x + y
	fmt.Println("z is -->",z)

	n1 := add(a, b)
	fmt.Println("z is -->",n1)

	getAbout("Suvo Datta",24,"Full Stack web Developer")
	getAbout("Suvo Datta",24,"Back End Developer")
	getAbout("Suvo Datta",24,"Front End Developer")

}