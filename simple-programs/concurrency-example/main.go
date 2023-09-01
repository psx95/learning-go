package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	demoVariant := flag.String("variant", "wg", "Which concurrency demo to run")
	flag.Parse()

	switch *demoVariant {
	case "wg":
		demonstrateConcurrencyThroughWaitGroups()
	case "ch":
		demonstrateConcurrencyWithChannels()
	default:
		fmt.Println(fmt.Errorf("invalid argument: %s", *demoVariant))
	}
}

func demonstrateConcurrencyThroughWaitGroups() {
	// Declare a WaitGroup to manage the counter
	var wg sync.WaitGroup

	wg.Add(1) // We are about to initiate 1 goroutine
	// Define the goroutine - our concurrent task
	go func() {
		fmt.Println("This work is async")
		wg.Done() // This is called within the concurrent task
	}()

	fmt.Println("This is some syncronous work")
	// Wait till all goroutines registered with the Go scheduler to be completed
	wg.Wait()
}

func demonstrateConcurrencyWithChannels() {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)

	go func() {
		// send message to the channel
		ch <- "Sent Message"
		// No done here since we do not care when this message is sent
		// as long as we are waiting on the task that is listening to
		// this channel.
	}()

	// This is the async work that we care about and need to wait for
	go func() {
		// recieve the message from the channel
		msg := <-ch
		fmt.Printf("Received: %s\n", msg)
		// Once we have recieved the message, we can call done so that
		// the wait block is removed.
		wg.Done()
	}()
	// wait till we have received the message from the channel
	wg.Wait()
}
