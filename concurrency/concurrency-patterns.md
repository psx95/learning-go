### Common concurrency patterns in Go
These are common patterns that emerge when writing concurrent programs in Go. These patterns help maintain concurrent code overtime, but they may not be applicable to all applications.
There are three main categories of concurrency patterns discussed here - 
1. Non-blocking error channels
2. Encapsulating goroutines
3. Foundational concurrency patterns

#### Non-blocking error channels
This patterns involves making sure that if there is an *error channel* used in our Go program, then it always has a valid reciever so that it does not block the execution of program in case
the consumer of the goroutine does not implement a receiver. 
```go
var (
    inCh = make(chan string) // input channel for recieving inputs
    outCh = make(chan int) // output channel to send the results
    errCh = make(chan error, 1) // error channel to log the errors
)
// Note that the error channel is actually a buffered channel with a buffered value
// of 1. This assures that the error channel always has a reciever.
// This makes the error channel a non-blocking channel. If this is not done, then it
// means that it is upto the consumer of the channel to drain the channel. 

// worker converts strings from input channel to integers and sends
// the results to an out channel. 
func worker(inCh <-chan string, outCh chan<- int, errCh<- error) {
    for msg := range in {
        i, err := strconv.Atoi(msg)
        if err != nil {
            // error occurred, cannot continue
            // report error to the error channel and return
            // If the error channel was unbuffered, then any errors 
            // reported to the channel should be consumed immediately,
            // else this call would have become a blocking one. 
            errCh<- err
            return
        }
        outCh <- i // success report converted int to the out channel
    }
}
``` 
*The buffered channel prvents a resource leak in case the consumer of the above function has not created a valid reciever for the error channel. 
With buffered channels, even if the user does not create a valid reciever to drain the channel, the garbage collecter will clean up channel along with any messages stored in its buffer.*

>NOTE: It could be said that the same case can happen if the user does not create a valid receiver for the out channel, but it is assumed that this would not be the case since it is primarily the way to retrieve the output from this function.

#### Encapsulating goroutines
