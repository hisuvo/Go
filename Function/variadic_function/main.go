package main

import "fmt"

//varidic function
func process(value ...int) {
	fmt.Println(value)
}

//argument in function
func greet(prefix string, mps ...string){
	// for idx, mp := range mps{
	// 	println(idx, prefix, mp)
	// }

	for _, mp := range mps{
		println(prefix, mp)
	}
}

func main() {
	process(2,3,4)
	// greet("Welcome","suvo","Anupom","Sajeeb") 
	mps := []string{"Suvo Datta","Aroshi Roy","Rajib Datta","Kanak Datta"}
	greet("Welcome to", mps...)
}