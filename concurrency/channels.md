### Channels in Go
Channels provide a reliable mechanism to coordinate the goroutines created in the program. Go runtime and Go scheduler schedules the available goroutines according to their own internal logic - this leads to a need for channels to coordinate how these goroutines are run.

Channels are types like anything in Go and therefore they can be stored as variables. However, channel is the only type in Go that must be built using the builtin `make` function.
A channel only allows passing one type of data through it and this type has to be specified when creating the channel. A channel allows two types of operations - 
1. Send a message into the channel
2. Receive a message from the channel

The following example shows how channels are used in Go - 
```go
// creating a channel of strings
// chan keyword specifies a channel
// a channel must be created using the builtin make keyword
ch := make(chan string)

// sending a message to the channel
// Operations on the channel are defined using the left-arrow operator (<-)
// The following line sends a string message into the channel
ch <- "hello"

// receiving a message from the channel
// This is also defined using the left-arrow operator (<-)
// The following line receives a message from the channel
msg := <- ch
```

**NOTE**
Channel operations block until complementary operations are ready. In other words, 
 - Sending operation will block at the sender until the receiver is ready.
 - Receiving operation will block at the receiver until the sender is ready.

This blocking behavior is what allows channels to synchronize goroutines and help them communicate with each other.

#### Looping with Channels
We can appply the concept of looping to channels to allow sending multiple messages to the same channel. In Go, this can be achieved by using a variation of the for loop.
One thing to note about using for loops with channels is that channels are a special type of collection - unlike other collections that Go has (slices, arrays, maps, etc.), a channel is not a closed closed collection.
This means that the size of the channel is unknown and there is always the possibility of having another message when looping over the channel.

So, to indicate that a channel will not have any more messages, we can use the `close(channel)` function. After this function is invoked, no more messages are allowed to be sent to the channel. 
Invoking this function also signals the for loops in the program that the channel does not have any more messages and this is how the loop gets closed.

The following examples illustrates how for loops can be used to loop through channels receive multiple messages.
 
```go
// Create channel
ch := make(chan int)

// goroutine that generates messages into the channel
go func(){
    for i := 0; i < 10; i++ {
        // put integer message into the channel
        ch <- i
    }
    // close the channel - so that no more messages are allowed to be sent
    // this signals the loops that are looping over this channel that there are no more messages and therefore the loop can be stopped.
    close(ch)
}()

// iterate through the channel using for - 
// each iteration pulls in a single message from the channel
for msg := range ch {
    // code to consume messages
}
```

#### Select statements
Select statement in Go are similar to switch statements, but are optimized to work with channels instead. They contain cases just like switches and are used to organize results coming back from multiple goroutines from various different channels.

The cases contain channel operations which can be either message send or message receive. Typically the `select` statement is a blocking operation and Go will not continue the execution of the program until one of the cases gets satisfied. To make it unblocking, we can add a `default` case, which is just like switch - executes when none of the other defined cases match.

One stark difference between select statement and switch statements is that in select, if more than one case can be acted upon, then one case is chosen at random. Unlike switch where the case defined first gets preference. 

The following code sample shows how select statement are used in a Go program -
```go
// Select statement works for channels
// Unlike switch, there is no condition or variable to evaluate
select {
    // Case for if we received a message from channel 1
    case msg := <- ch1:
        fmt.Println(msg)
    // Case for if we received a message from channel 2
    case msg := <- ch2:
        fmt.Println(msg)
}

// Without a default case, this is a blocking call and Go will not proceed until we receive a message 
// from either ch1 or ch2.
// If no goroutines are scheduled to run and the execution flow comes across a select statement which
// has cases that are waiting on some channels, Go issues a panic due to a deadlock - since with no goroutines 
// scheduled, there can be no way the program completes naturally. 
// This can be avoided by adding a default case.
```
#### Buffered and unbuffered channels

##### Unbuffered channels
So far, the channels shown in the above code snippets are examples of unbuffered channels. Unbuffered channels do not have a buffer capacity to hold the messages. In other words, when sending a message in an unbuffered channel, the go scheduler will block the execution of the go code untill there is a reciever to consume the message from the channel.

The code listening for messages on the channel is blocked as well till there is a message available in the channel and as soon as there is a message available, the flow on both - the sender and the receiver side is unblocked.

##### Buffered channels
Buffered channels have an internal buffer which is capable of holding a certain number of messages untill they are ready to be consumed by another receiver channel. This internal buffer prevents blocking the execution flow of the program since the channel itself acts like a receiver which immediately consumes the message. 

For instance, if a channel has a buffer size of 1 - this means that the channel is capable of holding 1 message in its internal buffer. This means, if a message is sent onto this channel, that message can be held in its internal buffer which can later be consumed by an external receiver. However, since the message is temporarily held in a buffer, the execution flow does not block and is therefore allowed to continue execution. 

```go
// unbuffered channel 
var ch = make(chan string)

func send() {
    // send message
    ch  <- "message" // the control flow is blocked on this line till a reciever consumes message
}

func receive() {
    // consume message from channel
    msg := <-ch // the control flow is blocked on this line till a sender sends a message
    fmt.Println(msg)
}

func main() {
    // syncronization code
    // ...
    go send()
    go receive()
}

// buffered channels
var ch_b = make(chan string, 1) // the second parameter is the buffer capacity for messages

// If ch_b is used in the send function, the send() function will not block till the reciever is ready,
// instead, the send function immmediately exists since the channel can hold 1 message internally in its 
// internal buffer.
// On the receiver side, things work similarly, since there is already a message on the internal buffer, 
// the gorouting receives the messages from the buffer and completes.
// This decouples the sending side of the channel from the recieving side of the channel.
```
