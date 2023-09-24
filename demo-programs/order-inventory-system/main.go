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
	// anonymous function to listen to valid order channel
	go func(validOrderCh <-chan order.Order) {
		validOrder := <-validOrderCh
		fmt.Printf("Valid order received: %v\n", validOrder)
		wg.Done()
	}(validOrderCh)
	go func(invalidOrderCh <-chan order.InvalidOrder) {
		invalidOrder := <-invalidOrderCh
		fmt.Printf("invalid order received: %v\n", invalidOrder)
		wg.Done()
	}(invalidOrderCh)
	wg.Wait()
}

func recieveOrders(outChannel chan<- order.Order) {
	for _, rawOrder := range rawOrders {
		var newOrder order.Order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("sending order to receieve order channel")
		// send order to outChannel for validation
		outChannel <- newOrder
	}
}

func validateOrders(inChannel <-chan order.Order, outChannel chan<- order.Order, errorChannel chan<- order.InvalidOrder) {
	incomingOrder := <-inChannel
	if incomingOrder.Quantity <= 0 {
		invalidOrder := order.InvalidOrder{
			Order:      incomingOrder,
			InvalidErr: fmt.Errorf("invalid order quantity: %v, order quantity should be greater than 0", incomingOrder.Quantity),
		}
		errorChannel <- invalidOrder
	} else {
		outChannel <- incomingOrder
	}
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": -42.5, "status": 1}`,
	`{"productCode": 2222, "quantity": 54.22, "status": 2}`,
	`{"productCode": 3333, "quantity": 61.23, "status": 3}`,
	`{"productCode": 4444, "quantity": 12.56, "status": 0}`,
}
