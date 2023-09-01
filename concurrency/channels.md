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
