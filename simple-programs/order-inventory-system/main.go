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
	wg.Add(1)
	go recieveOrders(&wg)
	wg.Wait()
	fmt.Println(order.Orders)
}

func recieveOrders(wg *sync.WaitGroup) {
	for _, rawOrder := range rawOrders {
		var newOrder order.Order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Println(err)
			continue
		}
		order.Orders = append(order.Orders, newOrder)
	}
	wg.Done()
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 42.5, "status": 1}`,
	`{"productCode": 2222, "quantity": 54.22, "status": 2}`,
	`{"productCode": 3333, "quantity": 61.23, "status": 3}`,
	`{"productCode": 4444, "quantity": 12.56, "status": 0}`,
}
