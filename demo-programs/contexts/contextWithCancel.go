package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	// WithCancel returns a cancel function that can be called once the work
	// associated with the derived context is complete.
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done() // call done when this function exists

		// this loops runs infinitely until the context is cancelled
		for range time.Tick(500 * time.Millisecond) {
			if ctx.Err() != nil {
				// if the Err() method returns a non-nil value, that means
				// the context has been cancelled
				log.Println(ctx.Err())
				return
			}
			fmt.Println("tick!")
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel() // cancel the context after 2 seconds

	wg.Wait()
}
