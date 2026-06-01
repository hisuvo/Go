package main

import (
	"fmt"
	test "recevier_test/Test"
)

type Teacher struct {
	name    string
	subject string
	perCls  int
	id int
}

// func printTeacher(teacher Teacher){
// 	fmt.Println(teacher.name)
// 	fmt.Println(teacher.subject)
// }

// Receiver function and It possible only for custome type 
func (teacher Teacher) printTeacher(){
	fmt.Println(teacher.name)
	fmt.Println(teacher.subject)
}

func (teacher Teacher) callTeacher(clsNo int){
	fmt.Println(teacher.name,"class Number ->", clsNo)
}

func Receiver(){
	var teacher1 Teacher

	teacher1 = Teacher{
		name: "Nirmolido",
		subject: "Argeculter",
		perCls: 4,
		id: 12,
	}

	// printTeacher(teacher1)

	// Receiver function call
	teacher1.printTeacher()
	teacher1.callTeacher(3)
}

func main() {
	// Receiver()
	test.TestOne()
}

func init() {
	fmt.Println("Receiver Function")
}