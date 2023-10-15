package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	// create a context that auto-cancels after 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// the context auto cancels, but we need to make sure that cancel runs
	// at least once and as early as possible (so if the operations associated with
	// context are finished before the deadline)
	defer cancel() // here it doesn't make much difference since cancel is called before main ends
	// NOTE: cancel function can be called more than once, but should be called at least once

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		// run loop every 500 ms infinitely
		for range time.Tick(500 * time.Millisecond) {
			if ctx.Err() != nil {
				// the context associated with this goroutine expired
				// NOTE: this is the only way the infinite loop can end in the
				// current program
				log.Println(ctx.Err())
				return
			}
			fmt.Println("tick!")
		}
	}(ctx)

	wg.Wait() // wait for the goroutine to finish
}
