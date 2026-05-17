package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) showInfo() {
	p.name = "Rajib"
	p.age = 18

	fmt.Printf("From ShowInfo Function: My name is %v and I am %v years old \n", p.name, p.age)
}

func (p *Person) changeName(){
	p.name = "Anupom Datta"
	p.age = 25

	fmt.Printf("From ChangeName Function: My name is %v and I am %v years old \n", p.name, p.age)
}

func main() {
	user := Person{
		name: "suvo",
		age:  24,
	}

	user.showInfo()
	user.changeName()
	fmt.Printf("From Main Function: My name is %v and I am %v years old", user.name, user.age)

}