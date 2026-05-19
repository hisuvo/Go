package main

import "fmt"

func multiplyBy(fector int) func(int) int {
	return func(i int) int {
		return i * fector
	}
}

func main() {
	dubble := multiplyBy(2)
	fmt.Println(dubble(5))
	fmt.Println(dubble(150))
}