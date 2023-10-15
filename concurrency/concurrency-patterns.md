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
    inCh = make(chan string) // input channel for receiving inputs
    outCh = make(chan int) // output channel to send the results
    errCh = make(chan error, 1) // error channel to log the errors
)
// Note that the error channel is actually a buffered channel with a buffered value
// of 1. This assures that the error channel always has a receiver.
// This makes the error channel a non-blocking channel. If this is not done, then it
// means that it is up to the consumer of the channel to drain the channel. 

// worker converts strings from input channel to integers and sends
// the results to an out channel. 
// This function is meant to be called as a goroutine.
func worker(inCh <-chan string, outCh chan<- int, errCh<- error) {
    for msg := range in {
        i, err := strconv.Atoi(msg)
        if err != nil {
            // error occurred, cannot continue
            // report error to the error channel and return
            // If the error channel was un-buffered, then any errors 
            // reported to the channel should be consumed immediately,
            // else this call would have become a blocking one. 
            errCh<- err
            return
        }
        outCh <- i // success report converted int to the out channel
    }
}
``` 
*The buffered channel prevents a resource leak in case the consumer of the above function has not created a valid receiver for the error channel. 
With buffered channels, even if the user does not create a valid receiver to drain the channel, the garbage collector will clean up channel along with any messages stored in its buffer.*

>NOTE: It could be said that the same case can happen if the user does not create a valid receiver for the out channel, but it is assumed that this would not be the case since it is primarily the way to retrieve the output from this function.

#### Encapsulating goroutines
This pattern helps guard concurrent Go programs against panics caused by `nil` channels. Sending a message to a nil channel causes a Go program to panic. 

We can obviously perform a nil check on a channel before sending a message into it, but an alternative to this is encapsulating the goroutine in a synchronous function itself. The synchronous function is responsible for creating the required output channels. The following example demonstrates this - 

```go
// The following snippet is the same example as used above.
// ORIGINAL SAMPLE
var (
    inCh = make(chan string) 
    outCh = make(chan int) 
    errCh = make(chan error, 1) 
)

// If the outCh and/or errCh passed by the user is nil or the channels are uninitialized,
// then in that case, this worker function will panic 
func worker(inCh <-chan string, outCh chan<- int, errCh<- error) {
    for msg := range in {
        i, err := strconv.Atoi(msg)
        if err != nil {            
            errCh<- err
            return
        }
        outCh <- i // success report converted int to the out channel
    }
}
// ORIGINAL SAMPLE

// The above code sample can be re-written to encapsulate the worker function logic in another synchronous function
// which is also responsible for creating and returning the output channels which can then be used by the consumers.
var inCh = make(chan string)

// This function is meant to be called synchronously
func worker(inCh <-chan string) (outCh chan<- int, errCh<- error) {
    // This function makes sure that the output channels are always valid
    outCh := make(chan int)
    errCh := make(chan error)

    // the core logic of the asynchronous work is written in an encapsulated go routine
    go func() {
        for msg := range inCh {
            i, err := strconv.Atoi(msg)
            if err != nil {            
                errCh<- err
                return
            }
            outCh <- i // success report converted int to the out channel
        }
    }()
    // these channels can now be used by the consumer to retrieve results from the goroutine 
    // and are assured to be valid.
    return outCh, errCh
}
```

#### Foundational concurrency patterns
There are four foundational concurrency patterns commonly found in Go and most other languages - 
1. Single producer, single consumer
2. Single producer, multiple consumers
3. Multiple producers, single consumer
4. Multiple producers, multiple consumers

 - **Producer(s)** is/are the goroutine(s) that is/are producing the message(s) in a channel.
 - **Consumer(s)** is/are the goroutine(s) that is/are recieving the message(s) from a channel.
