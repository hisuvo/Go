package main

import "fmt"

type Persion struct{
	name string
	age int
}

// this is reciver function
func (p Persion) introduce() {
    fmt.Println("My name is", p.name)
    fmt.Println("My age is", p.age)
}

func main() {
	user1 := Persion{
		name: "suvo",
		age: 24,
	}
	user2 := Persion{
		name: "Rajib Datta",
		age: 19,
	}

	user1.introduce()
	user2.introduce()
}