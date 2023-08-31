### Errors in Go
Errors in Go are unexpected situation that may arise when running an application. It is reasonable to expect errors especially when running an application at scale - for instace, a web server serving millions of requests. 
"Errors are Values" - in Go, errors should be treated as legitimate values returned from a function and should therefore not be ignored. Instead, it is better if these errors are handled immediately. Handling errors immediately leads to a more stable production code and increases error visibility in the program.

#### Error handling
Reminder: `error` in Go is a builtin type which is implemented as an interface with a single method named `Error()` that returns a string.

This means, we can create custom errors in Go by simply implementing the `Error() string` method in our custom type which can then be used as an error.
This also shows that `error` is a very lightweight object.

*NOTE: The `errors` package in Go help us to create new errors*

```go
// Creating errors
import (
    "errors"
    "fmt"
)

func main() {
    // Creating error objects in Go using the errors package.
    err := errors.New("My error")
    fmt.Println(err)

    // Creating error objects in Go using thr fmt package.
    // the formatting verb %w can be used to wrap one error into another
    err2 := fmt.Errorf("Wrap the first error with this one: %w". err)
    fmt.Prinltn(err2) // Wrap the first error with this one: My error
}
```

#### Error vs Panic
| Error  | Panic |
| -------| ------ |
| Error is basically the result of an operation, this means it is handled as part of general program flow. | Panic causes the function's execution to immediately terminate and it is propagated up the call stack until it sees a `recover` method. This means it causes a major change in the control flow of the program. |
| It is easy to discover an error since it is mentioned as a return value. By convention, it is always the last returned value | It relies on the programmer to read code and documentation to figure out if a function can panic |
| Error implies unexpected situation in a program | Panic implies that program is unstable (undefined behavior) |
| Errors are used more frequently | Panics are rare |
