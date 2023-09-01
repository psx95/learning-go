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
