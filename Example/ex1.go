package main

import "fmt"

func main() {
	s :=[]int{5, 5, 6, 7 , 8}
	s1 := []int{13,24,36}

	newSlice := append(s,s1...)
	fmt.Println(newSlice)
}

func init() {
	fmt.Println("Slice operation")
}