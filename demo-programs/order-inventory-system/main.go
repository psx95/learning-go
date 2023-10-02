package main

import (
	"encoding/json"
	"fmt"
	"log"
	"order-inventory-system/order"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	recieveOrderCh := recieveOrders()
	validOrderCh, invalidOrderCh := validateOrders(recieveOrderCh)
	reservedInventoryCh := reserveInventory(validOrderCh)

	// Add 2 to the counter, one for invalid orders, the other for inventory orders
	wg.Add(2)

	// consume all the invalid orders
	go func(invlidOrderCh <-chan order.InvalidOrder) {
		for invalidOrder := range invalidOrderCh {
			fmt.Printf("invalid order received: %v\n", invalidOrder)
		}
		wg.Done()
	}(invalidOrderCh)

	go func(reservedInventoryCh <-chan order.Order) {
		for reservedOrder := range reservedInventoryCh {
			fmt.Printf("Reserved Order: %v\n", reservedOrder)
		}
		wg.Done()
	}(reservedInventoryCh)

	wg.Wait() // wait till all orders are processed
}

func recieveOrders() <-chan order.Order {
	// outChannel could be passed as a parameter, but then
	// it would be necessary to guard against a nil channel
	// since it would lead to a panic.
	outChannel := make(chan order.Order)
	go func() {
		for _, rawOrder := range rawOrders {
			var newOrder order.Order
			err := json.Unmarshal([]byte(rawOrder), &newOrder)
			if err != nil {
				log.Println(err)
				continue
			}
			// send order to outChannel for validation
			outChannel <- newOrder
		}
		// close the outgoing channel - since all orders are received
		close(outChannel)
	}()
	return outChannel
}

// reserveInventory updates status of all valid orders received by the application to reserved.
func reserveInventory(in <-chan order.Order) <-chan order.Order {
	out := make(chan order.Order)
	go func() {
		// go through all valid orders received and store them in the inventory
		for o := range in {
			o.Status = order.Reserved
			out <- o // send updated orders to a channel
		}
		close(out)
	}()
	return out
}

func validateOrders(inChannel <-chan order.Order) (<-chan order.Order, <-chan order.InvalidOrder) {
	outChannel := make(chan order.Order)
	errorChannel := make(chan order.InvalidOrder)
	go func() {
		for incomingOrder := range inChannel {
			if incomingOrder.Quantity <= 0 {
				invalidOrder := order.InvalidOrder{
					Order:      incomingOrder,
					InvalidErr: fmt.Errorf("invalid order quantity: %v, order quantity should be greater than 0", incomingOrder.Quantity),
				}
				errorChannel <- invalidOrder
			} else {
				outChannel <- incomingOrder
			}
		} // will exit loop when inChannel is closed
		// once input channel is closed, close the outChannel and the errorChannel
		// since no more incoming orders to validate.
		close(outChannel)
		close(errorChannel)
	}()
	return outChannel, errorChannel
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": -42.5, "status": 1}`,
	`{"productCode": 2222, "quantity": 54.22, "status": 2}`,
	`{"productCode": 3333, "quantity": 61.23, "status": 3}`,
	`{"productCode": 4444, "quantity": 12.56, "status": 0}`,
}
