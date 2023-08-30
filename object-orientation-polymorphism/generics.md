### Generic programming in Go

Generic programming in Go is a way to overcome the limitations of interfaces - in the sense that whenever we use interfaces in Go, we loose the underlying concrete type with which we started.
Generic programming in Go offer us a way where we can use concrete types in generic functions, wherein they act *like interfaces* but once outside the generic function, we are free to use the
actual concrete type.

So taking the example of Go `File` (concrete type) and `io.Reader` (interface type), we could write a generic function, within which a `File` type could work as an interface type - (`io.Reader`),
but outside the function, the file object is retains its identity as a `File` type. 

Such a function is able to work on any `io.Reader` implementation, but does not cauase the concrete type to loose its concrete identity - at the end of the function a `File` object remains a `File`
and a `TCPConn` object (*a type from Go standard library which also implements the `io.Reader` interface*) would remain a `TCPConn` object. 

#### Creaating a generic function in Go -

```go
intArr := []int64{0, 1, 2}
floatArr := []float64{1.2, 2.1, 5,6}
// A normal function in Go 
// Foo takes in a slice of int64 and returns a slice of int64
func foo(arr []int64) []int64 {    
    result := make([]int64, len(arr))
    // function code to process arr
    return result
}

// Invoking foo - like a normal Go function
anotherArr := foo(intArr)


// Creating a generic version of foo that takes in a slice of type 'V'
// the '[V any]' after the function name in the signature is called the type constraint.
// In this function type constraint on type V is 'any', which means V can asssume any type
func bar[V any](arr []V) []V {
    result := make([]V, len(arr))
    // function code to process arr
    return result
}

// Invoking generic function bar - with a float64 type
floatAnotherArr := bar[float64](floatArr)

// Invoking generic function bar - with a int64 type
intAnotherArr := bar[int64](intArr)
```
