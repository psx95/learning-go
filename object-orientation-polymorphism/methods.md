### Methods in Go
Methods in Go work similar to how a function works in Go, but a method and a function are not the same thing. 
A method in Go has to have a custom type - a type defined and controlled by the programmer. The type can be 
anything. The following examples distinguish differences between a method and a function in Go -

```go
var i int 
// function
func isEven(num int) bool {
    return num%2 == 0
}

// method in Go
// First, we need to define a custom type on which the method can operate
// In Go, we can bind a type to anything - they don't necessarily have to be structs
// Here, we bind a custom type myInt to the builtin int
type myInt int
func (i myInt) isEven() bool {
    return int(i)%2 == 0
}

// NOTE: The (i myInt) between func keyword and the method name is called a method reciever
// The method receiver is the main difference between a method and a function in Go.
```
So, the main difference between a method and a function in Go is the **method reciever**. The method reciever indicates 
a tight coupling between a type and a function (the method) - the method will always execute in the context of a certain variable (in the above example, the variable i of type `myInt`).

From the above example, both the method and the function are performing the exact same thing, but we can pass any `int` to the function whereas the method can only be executed on a type `myInt`. 
This also causes a change in how a method is invoked vs how a function is invoked - 

```go
// Invoking the isEven function defined above. (We already declared i, so we can pass it)
ans := isEven(i)

// Invoking the isEven method defined above. 
// For this, we first need to declare/initialize an object of type `myInt` declared above.
var customInt myInt 
ans = customInt.isEven() // the method is invoked on the object here.
```

#### Method Receivers: Pointers vs Values

Consider the following example - 

```go
type user struct {
    id       int
    username string
}

// Method reciever acts on the value
func (user u) String() string {
    return fmt.Sprintf("%v (%v)\n", u.username, u.id)
}

// Method reciever acts on the the pointer
func  (user *u) UpdateName(newName string) {
    u.username = newName
}

// A pointer reciever is used to share a value between the caller and the receiver.
```

In the above example, for the value reciever since we do not want any side effects from calling the `String()` function on the `user` struct, we use a value reciever &rarr; this means we are not sharing data. 

However, in case of `UpdateName()` method, the intention is to update the username field in the `user` struct so we pass a pointer reciever to share data between the caller and the receiver.

Typically, it is better (more clear) to use value receivers whenever we are not trying to share data. 
