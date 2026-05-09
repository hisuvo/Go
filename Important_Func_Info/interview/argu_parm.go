package interview

import "fmt"

func about(name string, age int)  { // parameters
	fmt.Println(name,"is",age,"years old!")  
}

func GetInfo(){
	 about("Suvo Datta", 24) // arguments
}