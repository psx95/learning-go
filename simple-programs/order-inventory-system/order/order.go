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

const (
	none OrderStatus = iota
	received
	reserved
	filled
)

func orderStatusToString(status OrderStatus) string {
	switch status {
	case none:
		return "None"
	case received:
		return "Received"
	case reserved:
		return "Reserved"
	case filled:
		return "Filled"
	default:
		return "Unknown status"
	}
}
