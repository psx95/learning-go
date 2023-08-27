### Interfaces in Go

The concept of interfaces in Go is similar to what it is in other object oriented languages. It models certain behavior in objects and hides away the actual implementation details of that behavior *(which are defined in concrete types)*. 
This gives the programmer the flexibility to change the underlying implementation details, without affecting the rest of the application code that depends on that detail.

#### Defining and implementing interafces in Go

Interfaces are defined in a similar fashion to structs in Go - 

```go
// Defining interfaces in Go
// Reader interface helps to read binary data
type Reader interface {
    // Describe interface methods (behavior) here which all the implementing types are supposed to have
    Read([] byte) (int, error)
}

// Implementing interfaces
// Implementing interfaces doesn't necessarily have to be based on a struct
// File implements Reader -
type File struct { ... }
// File implements Reader interface defined above since it models all the behavior expected of a Reader
// We explicitly bind the Read method from Reader interface (uses same function signature) to the File
// type within the File type definition, which makes File implement Reader.
// NOTE: Interface implementation in Go is implicit, which means you do not have to explictly declare
// type A implments B (like in other object oriented language)
func (f File) Read(b []byte) (n int, err error)

// TCPConn implements Reader
type TCPConn struct { ... }
// Similar to File type above, all methods (behaviors) expected of Reader interface are bounded to TCPConn
// which makes TCPConn implement Reader 
func (t TCPConn) Read(b []byte) (n int, err error)

var f File  // Concrete type
var t TCPConn // Concrete type

var r Reader // Interface type

// Since both File and TCPConn implement Reader implicitly as explained above, Go allows us to assign concrete types to
// interface types. For eg - 

r = f // This is allowed

// This allows us to invoke any method which is part of the Reader interface on File type
r.Read(...) // invokes the functionality for Read method defined for the File type

r = t
r.Read(...) // invokes the functionality for Read method defined for the TCPConn type
```

#### Type Assertions - get concrete type from interface type

```go
// File implements Reader
type File struct { ... }
func (f File) Read(b []byte) (n int, err error)

var f File
var r Reader =  f   // As mentioned above, this works

var f2 File = r     // compile-time error, this is not allowed in Go

// Type assertion - asserting that a given variable is of a certain type
// This works here since we know the underlying type in r is actually a File\
f2 = r.(File)   // panics upon failure

// Type assetion - avoiding panic when underlying type is unknown
// the second parameter we accept here is a boolean, which indicates if the type
// assertion we made here was correct or not. 
//
// NOTE: In case the assertion fails here, then the value of f2 remains unchanged
// If f2 was never assigned anything prior to this, it remains null.
f2, ok := r.(File) // Does not panic

// TYPE SWITCH - making type assertions against multiple types -
// The example above only checks the variable against a single type - File
// In case we want to check for types against multiple types we can use type switch, eg -

switch v := r.(type) {
    case File:
        // v is now a File object
    case TCPConn:
        // v is now a TCPConn object
    default:
        // like normal swicth statenents, handle if the interface does not match any concrete type.
}
```
