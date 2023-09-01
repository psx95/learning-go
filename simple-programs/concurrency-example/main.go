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
	fmt.Println("Implement this")
}
