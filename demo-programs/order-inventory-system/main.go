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
	fillOrdersCh := fillOrders(reservedInventoryCh)

	// Add 2 to the counter, one for invalid orders, the other for filling orders
	wg.Add(2)

	// consume all the invalid orders
	go func(invlidOrderCh <-chan order.InvalidOrder) {
		for invalidOrder := range invalidOrderCh {
			fmt.Printf("invalid order received: %v\n", invalidOrder)
		}
		wg.Done()
	}(invalidOrderCh)

	go func(fillOrdersCh <-chan order.Order) {
		for filledOrder := range fillOrdersCh {
			fmt.Printf("Order Completed: %v\n", filledOrder)
		}
		wg.Done()
	}(fillOrdersCh)

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

// fillOrder completes all the reserved orders from the reserve inventory channel.
func fillOrders(in <-chan order.Order) <-chan order.Order {
	out := make(chan order.Order)
	var wg sync.WaitGroup

	const workers = 3
	wg.Add(workers)

	// multiple consumers of the same channel
	for i := 0; i < workers; i++ {
		go func(worker int) {
			for o := range in {
				fmt.Printf("order: %v, worker: %v\n", o.ProductCode, worker)
				o.Status = order.Filled
				out <- o
			}
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// reserveInventory updates status of all valid orders received by the application to reserved.
func reserveInventory(in <-chan order.Order) <-chan order.Order {
	out := make(chan order.Order)
	var wg sync.WaitGroup

	// add multiple producers for reserved inventory channel
	// multiple producers means spawning multiple workers to simultaneously push messages to the
	// reserved inventory channel
	const workers = 3
	// wait group that waits on all workers
	wg.Add(workers)
	// create multiple goroutines - each of them send messages to the same channel
	for i := 0; i < workers; i++ {
		go func() {
			// go through all valid orders received and store them in the inventory
			for o := range in {
				o.Status = order.Reserved
				out <- o // send updated orders to a channel
			}
			wg.Done()
		}()
	}

	// Another goroutine that closes the out channel at the appropriate time
	// The code in this goroutin could be executed syncronously as well, but that would block the main thread.
	go func() {
		wg.Wait() // wait till all workers in the wait group are completed
		// since all workers sending messages to the channel are finished, close the channel
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
