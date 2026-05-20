package order

type OrderStatus int

const (
	Pendding OrderStatus = iota
	Paid
	Shipped
	Delivered
	Cancelled
)
