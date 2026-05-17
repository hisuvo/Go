package main

import "fmt"

type Vehicle interface{
	start()
}

type Car struct{
	name string
	color string
	price int
	model string
}

type Bike struct{
	name string
	color string
	price int
	model string
}

type Bus struct{
	name string
	color string
	price int
	model string
}

func (c Car) start(){
	fmt.Println("Car name is", c.name, "and this price is",c.price,"tk")
}

func EngineStart(v Vehicle){
	v.start()
}

func main() {
	fmt.Println("Go interface")
	thair := Car{name:"Mahidra Thair",color: "Black",price: 1500000, model: "4x4"}
	fmt.Printf("Car name: %v, color: %v, pirce: %v, model: %v \n", thair.name, thair.color, thair.price, thair.model)
	EngineStart(thair)
}