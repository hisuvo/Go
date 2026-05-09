package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	subject string
	salary int
}

func main() {
	var person1 Person
	// var person2 Person
	var person3 Person

	// pserson1 specificitaion
	person1 = Person { // Instantiate
		name : "suvo",
		age: 24,
		job: "Backend Developer",
		salary: 3000,
	}

	// pserson2 specificitaion

	person2 := Person {   // Instance or Object
		name:"Rajib Datta",
		age: 18,
		job: "Student",
		salary: 3000,
	}

	person3.name ="Anupom Datta"
	person3.job = "Backend Developer"
	person3.subject = "Go Language"

	fmt.Println(person3)

	printPerson(person1)
	printPerson(person2)
}

func printPerson(person Person){
	fmt.Println(person)
	fmt.Println("My name is",person.name,"I am a ", person.job ," of Golang.")
}