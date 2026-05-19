package main

import (
	"fmt"

	"example.com/myapp/math"
	"example.com/myapp/payment"
	"example.com/myapp/utils"
)

func main() {
	sum := math.Add(10, 30)
	utils.Print("Hello go module split")
	payment.Pay();
	fmt.Println("sum:",sum)
}

func init(){
	fmt.Println("This is init function")
}