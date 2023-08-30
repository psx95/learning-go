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

#### Creating custom type constraints in Go -

There are only two builtin interfaces in go that can be used as type constraints for generics - `comparable` and `any`. This means that there maybe cases where we need custom constraint types. We can create custom type constraints in Go, which can then be used to define our generic functions -

```go
// Since type constraints for generic functions are basically interfaces, we can create a custom interface and use it as a constraint.
// However this interface would be a litte different from the behavioral interfaces that we typically create since rather that defining
// behavior on implementing types, such an interface will define the concrete types (basically the types that would satisfy the constraints) -

// Consider the following function that adds elements in an int slice using '+' operator
func add(arr []int64) int64 {
    var result int64
    for _, v := range arr {
        result += v
    }
    return result
}

// In case we wish to use the same function for ints, floats and strings - we cannot use 'any' or 'comparable' type -
// This is because the '+' operator is not defined on type 'any'
// So we create a custom interface

// contains types that have '+' operator defined on them
// This is a type interface - uses same structure as a typical behavioral interface but contains a list of concrete types
// separated by the '|' (pipe) operator.
type addable interface {
    int | float64 | string
}

// Add generic now has a type constraint 'addable' which means we can pass any of int, float64 or string to it.
func addGeneric[V addable](s []V) V {
    var result V
    for _, v := range arr {
        result += v
    }
    return result
}
```
