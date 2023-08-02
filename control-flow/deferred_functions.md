### Deferred functions in Go
Deferred functions in Go are a means to delay the execution of code. Deferred functions are executed right before the control flow is returned to the user. 

Deferred functions call are typically made from within a function, and are queued up to be executed right after the function that contains the deferred call is done executing. A simple example demonstrating this is - 

```go
func main() {
    fmt.Println("main 1")
    defer fmt.Println("defer 1")
    fmt.Println("main 2")
    defer fmt.Println("defer 2")
}

// So the output of executing the main function shown above would look something like - 

/**
main 1
main 2
defer 2
defer 1
**/

// This indicates that the deferred function calls are stored in a stack based data structure.
// Basically follows the Last In First Out order of execution.  
```

*The reverse order of execution for deferred functions is helpful since it's a common pateern in Go to acquire a closable resource such as a connection to a Database and place close call in a deferred function right next to it. This way the resources are closed in the reverse order of which they are acquired.*
