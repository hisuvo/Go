package main

import "fmt"

func forLoop(){
	fruits := [4]string{"Apple", "Banana", "Mango", "Orange"}

	for i:=0; i < len(fruits); i++{
		fmt.Printf("value is %v\n",fruits[i])
	}

	// modernized for loop
	for i:=range len(fruits){
		fmt.Printf("Modern Wasy get value is %v\n",fruits[i])
	}

	for _,value := range fruits {
		fmt.Printf("Value is -> %v and this type is -> %T\n", value, value)
	}
}

// while style loop
func whileLoop() {
	i := 1;
	for i < 5{
		fmt.Println("Loop number ->", i)
		i++
	}
}

// range use for looping
func rangeLoop(){
	 fruits := [3]string{"apple", "orange", "banana"}

  for idx, value := range fruits {
     fmt.Printf("%v\t%v\n", idx, value)
  }
}

func main() {
	// forLoop()
	// whileLoop()
	rangeLoop()
}

