package main

import (
	"encoding/json"
	"fmt"
	"log"
	"order-inventory-system/order"
	"sync"
)

func main() {
	recieveOrderCh := make(chan order.Order)
	validOrderCh := make(chan order.Order)
	invalidOrderCh := make(chan order.InvalidOrder)

	var wg sync.WaitGroup
	go recieveOrders(recieveOrderCh)
	go validateOrders(recieveOrderCh, validOrderCh, invalidOrderCh)
	wg.Add(1)

	go func(validOrderCh <-chan order.Order, invalidOrderCh <-chan order.InvalidOrder) {
	loop:
		// Keep running select until all orders are routed to the correct channels
		for {
			// select is blocking, we typically do not want to block main thread,
			// so move this select in a goroutine.
			select {
			case validOrder, ok := <-validOrderCh:
				if ok {
					fmt.Printf("Valid order received: %v\n", validOrder)
				} else {
					// this will break out of the outer-for loop
					// without the label, the break would have broken out of
					// inner select statement
					break loop
				}
			case invalidOrder, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("invalid order received: %v\n", invalidOrder)
				} else {
					// this will break out of the outer-for loop
					// without the label, the break would have broken out of
					// inner select statement
					break loop
				}
			}
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)
	wg.Wait() // wait till all orders are processed
}

func recieveOrders(outChannel chan<- order.Order) {
	for _, rawOrder := range rawOrders {
		var newOrder order.Order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("sending order to receive order channel")
		// send order to outChannel for validation
		outChannel <- newOrder
	}
	// close the outgoing channel - since all orders are received
	close(outChannel)
}

func validateOrders(inChannel <-chan order.Order, outChannel chan<- order.Order, errorChannel chan<- order.InvalidOrder) {
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
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": -42.5, "status": 1}`,
	`{"productCode": 2222, "quantity": 54.22, "status": 2}`,
	`{"productCode": 3333, "quantity": 61.23, "status": 3}`,
	`{"productCode": 4444, "quantity": 12.56, "status": 0}`,
}
