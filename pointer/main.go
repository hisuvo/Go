package main

import "fmt"

// Note:
// 1. pase by value
// use value copy and pest

// 2. pase by reference
// -> & ambersend or address of value
// -> * value at address

/*

? ** Go Code Running step **

main.go
   │
   ▼
Go Compiler
   │
   ▼
Binary File (.exe)
   │
   ▼
Operating System
   │
   ▼
RAM
   │
   ├── Code Segment (Machine Code)
   ├── Data Segment (Global Variables)
   ├── Stack (Local Variables)
   └── Heap (Dynamic Objects)

? ** Full Flow **
* প্রথমে Go source code compile হয়ে binary (machine code) তৈরি হয়। তারপর program run করলে Operating System সেই binary-কে RAM-এ load করে এবং Code Segment, Data Segment, Stack, Heap তৈরি করে।
*/


func print(numbers [] int){ // pase infinite value
	fmt.Println(numbers)
}

func printStr(fruits *[3] string){ // pase only three value
	fmt.Println(fruits)
}

type Student struct {
	name string
	age int
	salary float64
}

func main() {
	// pointer or address of memory (ram)
	number1 := 20

	// ambersand & => address of number1 | 824,633,825,000
	address := &number1 
	
	// value assign at the number1 address
	*address = 50
	
	fmt.Println("Address of number1", address) // memory address of address
	// * => value at address
	fmt.Println("Value the address of number1 ->", *address) // 50

	arr := [] int{2,3,4,34, 45, 56, 60} // pase infinite value that a slice
	print(arr) // [2,3,4....., 60]
	
	arr1 := [3] string {"apple", "mango","Orrange"} // pase only 3 value that a array
	printStr(&arr1) // &["apple", "mango","Orrange"]
	fmt.Println(&arr1) // &["apple", "mango","Orrange"]

	student1 := Student{
		name: "Suvo Datta",
		age: 24,
		salary: 233456,
	}

	stu := &student1

	fmt.Println("student ->",stu) // {Suvo Datta 24 233456}
}


/*
Program শুরু হলে কোন function প্রথম call হয়?
আসলে main() সরাসরি প্রথম function নয়।
Go Runtime আগে কিছু initialization code চালায়।

সরলভাবে flow:
OS
 │
 ▼
Go Runtime Start
 │
 ▼
init() functions (যদি থাকে)
 │
 ▼
main.main()
 │
 ▼
print()
 │
 ▼
printStr()
 │
 ▼
Program Exit

** bionary code convert **
1001100
1001100
1001100
1001100
1001100
1001100

 ** code segment **
 func print() {...}
 func printStr() {...}
 tye Student struct {...}
 func main() {...}

 ** data segment **

 ** stack **

*/