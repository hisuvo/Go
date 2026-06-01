package functions

import "fmt"

// First-Order Function
func sum(a, b int) int {
	return a + b
}

// Higher-Order Function
func calculate(operation func(int, int) int, x, y int) int {
	return operation(x, y)
}

// Closure Creator
func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func HigerOrderFunc() {

	// First-Class Function
	var myFunc func(int, int) int = sum

	// Callback Function
	result := calculate(myFunc, 10, 20)

	fmt.Println("Result:", result)

	// Closure
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}