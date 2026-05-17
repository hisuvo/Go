package main

import "fmt"

/*
	1. create slice from an existiong array
	2. create slice from another slice
	3. create slice using make() function
	4. using slice literal  // []datatype{values..}
	5. make function with len
	6. make function with len cap
	7. emty slice or nil slice
*/


func main() {

	// slice careate laterial
	s := []int{};
	s = append(s, 2)
	s = append(s, 34)

	x := s
	x = append(x, 243)
	s = append(s, 90)

	fmt.Println("s -->",s)
	fmt.Println("x -->",x)
	

	// // 1. create slice from an existiong array
	// //========================================

	// arr := [6]string{"This","is","a","Go","interview","question"}
	// // fmt.Println(arr)
	
	// // slice index 1 to 4 ["is", "a", "go"] len->3, cap-> 5
	// s := arr[1:4] 
	// fmt.Println("slice1 ->",s) // len(s) cap(s)

	// 2. create slice from another slice
	//=====================================

	// arr := [6]string{"This","is","a","Go","interview","question"}
	// s := arr[1:4]

	// s1 := s[1:3] 
	// fmt.Println(s1) // ["a","Go"]
	// fmt.Println(len(s1)) // len -> 2
	// fmt.Println(cap(s1)) // cap -> 4

	// ss := arr[5:6] // ["quesiont"]
	// fmt.Println("slice2 ->",ss)


	// 3. create slice using make() function with len
	// ========================================
	// If capacity parameter not define then, it will be equal to length
	// s := make([]int,3) // [0,0,0], len-> 3, cap -> 3
	// s[1] = 4 // [0,4,0]
	// fmt.Println("slice",s,"lent->",len(s),"cap->",cap(s))


	// 4. create slice using make() function with len
	// ========================================
	// s1 := make([]int,3, 5) // [0,0,0], len-> 3, cap -> 3
	// s1[1] = 4 // [0,4,0]
	// fmt.Println("slice",s1,"lent->",len(s1),"cap->",cap(s1))


	// 6. using slice literal  // []datatype{values..}
	// ========================================
	// slice := []int{1,2,3,4,5,6,7}
	// fmt.Println("slice",slice)
	// fmt.Println("len ->",len(slice))
	// fmt.Println("cap ->",cap(slice))


	// 7. emty slice or nile slice
	// var s []int  // ptr=nil, len=0, cap=0
	// fmt.Println(s)
	// fmt.Println(cap(s))
	// fmt.Println(len(s))

}
