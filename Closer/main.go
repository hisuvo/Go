package main

import "fmt"

const a = 45;
var p = 100;

func outer() func(){
	mony := 100
	age := 24

	fmt.Println("my age is", age)

	// If the variable’s lifetime ends inside the function →
	//  it stays in stack, else goes to heap.
	show := func () {
		mony = mony + a + p;
		fmt.Println(mony)
	}
	return show
}

func call(){
	int1 := outer()
	int1()
	int1()

	int2 := outer()
	int2()
	int2()
}

func main() {
	call()
}

func init(){
	fmt.Println("Hello world!")
}



/* 

Runtime Execution (Memory Layout)
When you run that binary (./main), the Go runtime loads it into memory like this:

+--------------------+
| Code Segment (Text)| → compiled instructions
+--------------------+
| Data Segment       | → global & static variables
+--------------------+
| Heap               | → dynamically allocated memory
+--------------------+
| Stack              | → function call frames
+--------------------+


Let’s break this down ⬇️

1. Code Segment
Contains machine code generated from your .go file.
Read-only.
Executed directly by the CPU.

2. Data Segment
Stores global and static data.
Example:
var counter int = 100

counter lives in data segment.
3. Heap
Stores dynamically allocated data (new, make, slices, maps, etc.).
Managed by Go’s garbage collector (GC).

4. Stack
Each Goroutine has its own stack (starts small, grows as needed).
Contains:
Function arguments
Local variables
Return addresses

⚡ Go Memory Management
Go has automatic memory management, which means:
You don’t manually allocate/free memory.
Go runtime’s garbage collector frees memory that’s no longer reachable.


*/