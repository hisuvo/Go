package main

import "fmt"

// Normal function
func example()int{
	result := 10

	defer func ()  {
		result += 200
		fmt.Println("defer inside result:", result) // 210
	}()

	fmt.Println("I am from example function", result) // 10

	result += 100

	return result // 110
}

// Name + defer
func deferExample()(result int){
	result = 10
	defer func ()  {
		result += 200
		fmt.Println("defer inside result:", result) // 210
	}()

	fmt.Println("I am from example function", result) // 10

	result += 100

	return result // 310
}

func main() {
	// result := deferExample()
	fmt.Println("Return result:",example())
}