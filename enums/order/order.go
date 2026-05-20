package order

import "fmt"

type Order struct {
	ID     int
	Status OrderStatus
}

func UpdateOrderStatus(order *Order, status OrderStatus) {
	order.Status = status
}

func PrintOrder(order Order){
	fmt.Println("Order Id:", order.ID)
	fmt.Println("Order Status",order.Status)
}