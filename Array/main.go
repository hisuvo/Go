package main

import "fmt"

func arrayBaisc () {
	letter := [4] string {"A","B","C","D"}
	number := [] int {2,3,4,5,6,7}
	fmt.Printf("letter %v number %v\n",letter,number)

	arr1 := [5] bool {}
	fmt.Println(arr1)

	// specifice initial array is possible to decleare
	arr2 := [4] int{1:50, 3:100}
	fmt.Println(arr2[3])
}

func result()  {

	var numbers = [5]int{}
	persons := [3]string{}
	allPersons := [...] string {"a",3:"b","c","d"} // inferred

	fmt.Println(numbers,persons,allPersons)


	var price = [3] int {10,20,30};
	price[1] = 60
	fmt.Println(len(price))
	fmt.Printf("%T",numbers)

	// fmt.Println(arrayBaisc())
}


func main() {
	arrayBaisc()
}