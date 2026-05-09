package main

import "fmt"

func Auth(age int) string {

	// this is expression
	isAdult := age > 18

	if isAdult {
		// code to be execute if condition is true
		return "You are alizable enter this course"
	} else {
		// code to be execute if condition is false
		return "You are not alizable enter this course"
	}

	//? else_if condition
	// use the else_if statement to specify new condition if first condition is false

}

func elseIf(number int) {
	if number > 10 {
		fmt.Println("x is larger than or equal to 10.")
	} 
	 if number > 20 {
		fmt.Println("x is larger than 20.")
	} else {
		fmt.Println("x is less than 10.")
	}
}

func serverDb() error {
	fmt.Println("Database connected")
	return nil
}

func main() {
	// fmt.Println(Auth(16))
	elseIf(30)

	// here if statement inside x value assing so that outside i can use x
	if x := 30; x > 20{
		fmt.Println("30 is getter than 20")
	}

	if err := serverDb(); err != nil {
	fmt.Println("Database inside accoured error!")
}
}