package test

import "fmt"

type Student struct {
	Name    string
	Class   string
	Subject string
}

func (std *Student) sudentDetails(){
	fmt.Printf("memory address %p\n",&std.Subject)
}

func TestOne() {
	student1 := &Student{
		Name: "Anik",
		Class: "XI",
		Subject: "Mathematices",
	}

	student1.sudentDetails()
	
}