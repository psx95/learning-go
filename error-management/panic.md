### Panic and Recovery in Go
#### What is a Panic ?
Panics in Go represent a condition where a program is no longer able to execute in a reliable manner. Panics can come from either the code that we have written or the Go runtime.
When a panic occurs during the execution of an application, the entire application panics all the way up to the call-stack, until it reaches the Go runtime. 
In a basic panicking scenario, the Go runtime handles the panic by closing the application out and printing a message indicating what happened.

#### Recovery in Go
Recovery mechanism in Go works with deferred functions and contains instructions that bring the application back to a stable execution environment. This works because when a panic occurs, it immediately exits the function, which invokes the deferred function (which was registered before invoking the function that caused the panic) containing instructions on how to recover from the panic. 

*The end result being - due to the recovery mechanism through the deferred function call, the panic did not propagate up the call stack.*

The above concepts of panic and recovery are demonstrated in the following Go code - 

```go
func main() {
    fmt.Println("main 1")
    func1()
    fmt.Println("main 2")
}

// This function might panic, so we setup a recovery mechanism that helps to recover from that panic
func func1() {
    // Recovery mechanism registered via a deferred function
    defer func() {
        // recover() is a builtin function in Go, that receives whatever is passed in panic.
        // information from the passed argument can be used to actually recover from the panic.
        // a call to recover() basically prevents the calling function from being terminated due to panic.
        fmt.Println(recover())
    }()

    fmt.Println("func1 1")
    
    // panic() is a builtin function in Go
    // any argument can be passed to the panic function - this argument provides information
    // to the recovery method - which might be helpful to recover from the panic.
    // without any recovery mechanism registered, the application would have terminated at this panic.
    panic("Something went wrong") 
    
    fmt.Println("func1 2")
}
```

When the above code is executed, the following output is generated - 
```shell
main 1
func1 1
main 2

# We can see that panic caused func1 from being completed by terminating it, but the calling function - 
# main, did not terminate and was ran till completion - due to a deferred call to the recover() function.
```

It is important to note here that while `recover()` has a special function - preventing the panic from propagating up the call stack, the deferred function that invokes `recover()` is just a standard deferred function. 
As such, the deferred function will be invoked regardless of whether there was a panic situation or not. When the recover function is invoked without a panic situation, printing the result of `recover()` like we have 
done above, will result in `nil` value. This basically indicates that there was no panic.
