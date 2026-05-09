package main

import "fmt"

func coffeeOrder(customerName string, coffeeType string, Price int) string {
	getOrder := fmt.Sprintf("Order for %v: %v coffee costs %d taka",customerName, coffeeType, Price)
	return getOrder
}

func main() {
	 result := coffeeOrder("Shuvo", "Back", 150)
	 fmt.Println(result)
}
