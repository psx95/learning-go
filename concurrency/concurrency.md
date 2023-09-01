### Concurrency in Go
#### Concurrency vs Parallelism
Concurrency is not parallelism. Concurrency means that our application can do more than 1 thing at a time. It doesn't necessarily mean that it will. 

Parallelism is the act of executing more than one task at the same time. This means without concurrnecy, we cannot have parallelism.

Putting the above together, in Go, we can write concurrent programs, but we cannot force them to run in parallel. This is decided by the OS and the Go runtime environment.
The main difference between the concept of concurrency and parallelism is that parallelism implies that multiple tasks run simultaneously, at the same time. Concurrency means that while multiple tasks may be running simultaneously, only one of them would be actively worked on at a given time (basically keep switching between the tasks). 

For example, if we have 3 tasks - Task 1, 2, 3. 
 - b/w t=0 and t=5 work on Task 1
 - b/w t=5 and t=8 work on Task 2
 - b/w t=8 and t=10 work on Task 1 *again*
 - b/w t=10 and t=12 work on Task 3
 - b/w t=12 and t=15 work on Task 2 *again*

 So while there are 3 concurrent task (tasks running simultaneously), only 1 of them ever gets worked on at a given point in time. 

 In cotrast, if we have 3 tasks running parallely - Task 1, 2, 3, it would mean that all 3 of them are being worked on simultaneously.

#### CSP (Communicating Sequential Processes)
CSP is the idea of breaking the program into discrete execution units of work (or tasks) which some processes can execute sequentially. So basically one process does some work and then passes the result of that work to something called a `channel`. 
Another process receives that result from this channel and then does some additional work on it (or finishes it) and so on. So these processes are communicating with each other via channels in a sequential manner - which is the main idea behind CSP.

The use of this model is that the processes are independent of each other and can work concurrently as long as there is some syncronization mechanism in place that makes sure of the sequential communication. So CSP makes concurrency patterns like `fan-in`
and `fan-out` easy to implement.

In context of Go language in particular, the processes described above are called `goroutine`s which can communicate with each other via `channel`s.

#### Goroutines
A goroutine is basically a function that can execute concurrently with other goroutines defined in the same program. Goroutines are lightweight and costs just a little more than allocation on the stack space. 

These goroutines are executed by the Go scheduler which is a part of the Go runtime environment. The Go scheduler executes these routines on its own and the programmer does not have any control over when these goroutines would be executed.

The following code sample shows how goroutines are created -
```go
func main() {
    // Some code
    
    // This is a goroutine
    // It is merely an anonymouos inline function that has the keyword 'go' in before the 'func' keyword.
    // The 'go' keyword is responsible for creating goroutines, without the go keyword, this would just be
    // an inline anonymous function.
    go func() {
        fmt.Println("An async task scheduled by the Go scheduler")
    }()

    // NOTE: This goroutine in its current form here will not do anything. This is because declaring it here just
    // informs the Go scheduler that there is a concurrent task that needs to be performed, but the program gets
    // terminated on the next line, which means the scheduler did not have any time to execute this goroutine.
    // WaitGroups can help here - they are explained in the next section.
}
```


#### WaitGroups
WaitGroups are essentially counters which exhibit special behavior when their value reaches 0. Since it is a counter, we can - 
1. Increment it - `Add(delta int)`: Increment counter by the passed delta.
2. Decrement it - `Done()`: Decrement counter by 1.
3. Wait on it - `Wait()`: Blocks the execution of program till the counter value reaches 0.

Typically in Go, there will be a single function that creates all the concurrent tasks, calling `Add` with the number of created tasks and then wait via the `Wait()` method till the counter reaches back 0. Each of the concurrent tasks would call `Done()` when they are complete and so when the last task completes, the block introduced by `Wait()` is removed.

The following code sample showcases the use of WaitGroups - 
```go
func main() {
    // Declare a WaitGroup
    var wg sync.WaitGroup

    // Add a concurrent task (the goroutine we created just below this line represents a concurrent task)
    wg.Add(1)
    // Create a goroutine
    go func() {
        fmt.Println("Some async work is being done here")
        // Indicate that the concurrent task is now completed
        wg.Done() // Decrements the WaitGroup counter by 1
    }()
    
    // Some regular syncronous code 
    fmt.Println("This is some syncronous work")
    
    // Wait till all concurrent tasks are completed
    // As soon as the counter reaches 0, this call no longer blocks and the execution is resumed.
    wg.Wait()
}
```

#### Channels
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
