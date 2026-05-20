package main

import "enumsapp/order"

func main() {
	myOrder := order.Order{
		ID:     1,
		Status: order.Pendding,
	}
	order.UpdateOrderStatus(&myOrder,order.Pendding)

	order.PrintOrder(myOrder)
}