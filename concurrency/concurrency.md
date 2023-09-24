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

##### More details on Goroutines:

Goroutines represent the concurrent tasks tha we as programmers create and they 'sit' between the go program which is the code we write and the Go scheduler. In other words, this means that goroutines interact between our program and the scheduler.

The Go scheduler in-turn interacts with the underlying OS using constructs called threads - regardless of the OS being used. 
Threads are tools which the operating systems use to manage their concurrency.

Goroutines can be considered as virtual threads which are maintained purely within the Go program. 
Because Goroutines are maintained purely within the Go program, they are limited and cannot do much on their own.
The Go scheduler maps these goroutines onto actual OS threads which is where the concurrent tasks represented by goroutines is actually executed. In other words, if a goroutine is not scheduled onto an OS thread, it will not be able to do anything.

*NOTE: Many languages do not have this virtualization concept and so to perform concurrent tasks in those languages, one would need to create an OS thread directly and perform work on that thread.*

A comparison between a goroutine and a thread can be summarized in the table below :
| thread  | goroutine |
| -------| ------ |
| Have their own execution stack | Have their own execution stack |
| Fixed stack space (~1MB) | Variable stack space (~starts at 2KB, max upto ~ 2GB) |
| Managed by OS (slightly more efficient, since no scheduler overhead) | Managed by Go runtime |
| Relatively expensive (interactions between OS and the program) | Inexpensive since its managed by runtime, need only talk to the go runtime |

##### Lifecycle of a goroutine

Broadly speaking, there are 3 main stages in the lifecycle of a goroutine - 
1. Create
2. Execute
3. Exit

However, since goroutines are scheduled for execution by the scheduler, this means whenever a goroutine is not scheduled to run on a thread by the scheduler - it is in a *blocking* state. This simply means the goroutine while created, is not doing anything.
Whenever the goroutine is scheduled to run on a thread by the scheduler, go routine moves to a *running* state. In the execution phase, the goroutine may switch between these two states, depending on the kind of work it is doing. 

For instance, if a goroutine is waiting on some input from the user, the scheduler may unschedule it from an OS thread - causing the goroutine to move to a *blocking* state, once it recieves the input, the scheduler would move the goroutine back to the *running* state.

Some other examples of why a goroutine might move to a blocking state include: 
 - The goroutine maybe waiting on a system call. Eg - the goroutine might be creating a file, which is a system call, in this case the goroutine will have to wait till the system call is complete. 
 - The goroutine might simply be sleeping. Eg - the `time.sleep` function tells the goroutine to wait for a certain period of time before continuing the execution.
 - Waiting on a network call - the goroutine could make a network call and may have to wait for receiving a response before continuing. In this case there is no value for the goroutine to be scheduled onto a thread simply waiting for the response.

In each of the above cases, the scheduler will not schedule the goroutine to run on an OS thread - since there is no active work to be done. This also means that the go scheduler may pull off already scheduled/executing goroutines from the OS threads if there is no active work to be performed. The scheduler reschedules these goroutines once they are out of the blocking state by scheduling them onto OS threads.


**NOTE:** Goroutines also have the responsibility of executing deferred functions. If a goroutine has a deferred function that it initiates, then the goroutine becomes the execution context of the deferred function and it will execute the deferred function after completing its execution.

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
