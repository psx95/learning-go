package order

import "fmt"

type Order struct {
	ProductCode int
	Quantity    float64
	Status      OrderStatus
}

type InvalidOrder struct {
	Order      Order
	InvalidErr error
}

type OrderStatus int

func (o Order) String() string {
	return fmt.Sprintf("Product Code: %v, Quantity: %v, Status: %v\n", o.ProductCode, o.Quantity, orderStatusToString(o.Status))
}

func (io InvalidOrder) String() string {
	return fmt.Sprintf("Product Code: %v, Quantity: %v, Status: %v, Issue: %v\n", io.Order.ProductCode, io.Order.Quantity, orderStatusToString(io.Order.Status), io.InvalidErr.Error())
}

const (
	None OrderStatus = iota
	Received
	Reserved
	Filled
)

func orderStatusToString(status OrderStatus) string {
	switch status {
	case None:
		return "None"
	case Received:
		return "Received"
	case Reserved:
		return "Reserved"
	case Filled:
		return "Filled"
	default:
		return "Unknown status"
	}
}
