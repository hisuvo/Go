package main

import "fmt"

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("start")
	panic("crash happend")
}

func main() {
	test()
	fmt.Println("Program continu....")
}