package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	demoVariant := flag.String("variant", "wg", "Which concurrency demo to run")
	flag.Parse()

	switch *demoVariant {
	case "wg":
		demonstrateConcurrencyThroughWaitGroups()
	case "ch":
		demonstrateConcurrencyWithChannels()
	case "sel":
		demonstrateSelectStatements()
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

func demonstrateSelectStatements() {
	// create multiple channels
	ch1, ch2 := make(chan string), make(chan string)

	// create goroutines to communicate on channels
	go func() {
		ch1 <- "Message to channel 1"
	}()

	go func() {
		ch2 <- "Message to channel 2"
	}()

	// This is just to give Go scheduler some time to schedule both the routines
	time.Sleep(10 * time.Millisecond)

	// Since both cases are now valid, a case will be selected at random.
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		// Default will PROBABLY not be executed now since the Go scheduler has
		// had time to schedule and run (10 ms) at least one of the 2 goroutines.
		// If the sleep statement was not there, there was a HIGH CHANCE that this
		// case might have ran.
		// But there is no guarantee for the above.
		fmt.Println("No messages received")
	}
}
