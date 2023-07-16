### Variables in Go

#### Declaration and Initialization
Variabels in Go are declared using the keyword `var`. There are various ways to declare and initialize variables in Go - 

```go
var name string         // declare variable
var name string = "psx" // declare and initialize

var name = "psx"        // declare and initialize with inferred type
myNmae := "psx"         // shorthand for declaration and initialization with inferred type

a, b := 10, 5           // Go allows multiple variables to be initialized all at once

// The shorthand declaration syntax is used most often in Go programs.
```

##### Notes
1. In Go, having an unused local variable is a compile-time error, so if you declare a variable and do not use it in your program, the Go program will not compile.
2. When using inferred types in Go, where there are multiple valid types for a literal value, Go has default data types to which it can default to. For instance, consider:

    `floatingNumber := 3.1342` 

    In the above piece of code, there are 2 possible inferred data types - `float32` and `float64`. In case for floating point numbers, Go has the default data type - `float64`. 
3. When assigning multiple variables in the same line, the initialization occurs in order - so in the example shown above, `10` is assigned to `a` and `5` is assigned to `b`.


#### Type conversion
Go **does not support implicit conversions between data types***. If a type conversion is required, it has to be explicit. So for instance, 

```go
var i int = 5   // declare and initialize

var f float32   // declare
f = i           // ERROR - implicit conversion not allowed in Go and will result in an error.
f = float32(i)  // Have to be explicit with the intent to convert to a different type.
```

### Constants in Go
Constants in Go are declared using the `const` keyword. Constants cannot change values once they are initialized and the value of the constant has to be determinable at compile time. Unlike variables constants do not need to be declared locally and so we can have unused constants in a Go program without any compile-time errors.

```go
const a = 42                // implicitly typed constant 
const name string = "psx"  // explicitly typed constant

const c = a               // one constant can be assigned to another

// Constant block - a group of constants
const (
    d = true       // constant with literal value true
    e = 3.144      // constant with literal value 3.144
    f              // In a constant block, if no value is provided, the value is copied from above - so f has 3.144
)

const g = 2 * 5                 // constant expression
const h = "hello, " + "world"   // should be calculable at compile time
const i = fooFunction()         // this is not allowed, since it cannot be determined at compile time. Functions require memory allocation which happens at runtime.

// Special case in constants - iota
const j = iota                  // j is treated as a constant group of size 1 and iota starts at 0, so j is assigned 0
const (
    k = iota                    // k is assigned 0
    l                           // iota is copied down, but iota's value is now updated to 1 (relative to its position in constant group) so l = 1
    m                           // m is assigned 2
    n = iota * 3                // n is assigned 9 since iota is 3
)
const (
    o = 10
    p = 4.5
    q = iota                    // even though this is first usage of `iota` in this group, since it is at 3rd position in group, it's value is 2
    r = 5
    s = iota                    // s = 4
)
```

In case of implicitly typed constants in Go, Go treats the constants as a literal value, so in the example above, `a` will be treated as a literal value, so whereever we use `a` in the program, Go will simply copy the value `42` in its place. This means we could use `a` to assign a floating point value as well - which would not be possible in case of variables.

In constrast, in case of explicitly typed constant, like `name`, it can only ever be assigned where a `string` datatype is allowed. 

##### Notes
1. Regarding iota - `iota` is a special symbol in Go that is used in the context of assigning the constants. `iota` has a value that is related to its position in a constant group and `iota`'s value at the first position in a constant group is `0`. The value of `iota` resets between each constant block. 
