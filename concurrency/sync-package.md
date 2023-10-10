### The sync package
The sync package in Go contains some useful features that help in writing concurrent programs. This document contains information on a few such features that help in certain situations like mutexes and once. 

#### Mutex
Stands for *mutual exclusion*. Mutexes are used in Go to acquire *locks* on resources that are shared across multiple goroutines so that at any given time no two goroutines access the same shared resource. 

Only the goroutine that acquired the *lock* is able to continue, the scheduler blocks the other goroutines that attempt to acquire the lock until the first goroutine releases the acquired lock. This becomes more clear through the example shown below -

```go
// some go file
var m sync.Mutex // goroutines can acquire locks on this resource

func main() {
    // goroutine #1
    go func() {
        m.Lock() // acquire lock on the mutex
        defer m.Unlock() // release the lock at the end of the goroutine
        // some application code that requires mutually exclusive access
        // ..
    }()

    // goroutine #2
    go func() {
        m.Lock() // attempts to acquire lock on the same mutex
        defer m.Unlock() // release lock at the end of the goroutine
        // some application code that requires mutually exclusive access
        // ..
    }()
}
```

Since both goroutines in the above code attempt to acquire locks on the same resource, only one of them will be granted access, the other will be blocked by the scheduler until it can acquire the lock. 

Assuming `goroutine #1` acquires lock first, `goroutine #2` will have to wait till `goroutine #1` releases the lock. Once it does release the acquired lock, `goroutine #2` then acquires it and is allowed to continue it's execution. 

>*NOTE: Mutexes are very good in protecting shared memory, but there is a performance penalty associated with their use, so in general shared memory should be avoided, but sometimes shared memory makes most sense for certain situations, in which case mutex can come in handy.*

#### Once
Once object in the sync package is used to ensure that something happens only one time.

This is commonly used in initialization logic in programs either because it can cause some unpredictable behavior or it would be inefficient to initialize multiple times.

```go
// Go program to open connection to a SQLite Database

func main() {
    foo()
    foo()
    // Once object guarantees that the function wrapped inside it is executed only once.
}

var db *sql.DB
var o sync.Once

// We would ideally want foo to be run only once, regardless of how many times
// it is invoked.
func foo() {
    // Once only has a single method on it named Do
    o.Do(func() {
        log.Println("opening connection to Database")
        var err error
        db, err = sql.Open("sqlite3", "./mydb.db") // driver name and connection string
        if err != nil {
            log.Fatal(err)
        }
    }) // this will run only once.
    // ...
    // ...
    // other code inside foo will still be run as many times as foo is invoked.
}
```
