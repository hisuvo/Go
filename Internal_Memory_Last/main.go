package main

import "fmt"

const birthYear int = 2001
var number int = 45

func call(){
   add := func (num1 int, num2 int)  {
      total := num1 + num2
      fmt.Println(total)
   }

   add(number, 3)
   add(50,40)
}

func main() {
   call();
	fmt.Println(number)
}

func init() {
   fmt.Println("My Birth Year Is ->",birthYear)
}

/*
  phase
   1.compilation phase
   2.execution phase

   ========= copiler phase ========
   *** code segment *** 
    birthYear = 2001
    call = func () {....}
    add = func () {....}
    main = func () {....}
    init = func () {....}

    *** data segemebnt ****
    number = 45


// For windows information given blow:   
   -> go run main.go => compile it => main.exe => main.exe
   -> go build main.go => compile it => main.exe
   -> main.exe
   
*/




/*

   Go Language code Internal Memory 
   i. code segment
   Read only code keep there like const, func, mian(), init()


   ii. data segment (Golobal Memory)
   Chanage able code keep there like assing value , var, := etc


   iii. Stack 
   here execution code take a memory this say stack frame


   iv.Heap
   This Task is GC (Gerbase collector)


*/