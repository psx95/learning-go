### Variables in Go

#### Declaration and Initialization
Variabels in Go are declared using the keyword `var`. There are various ways to declare and initialize variables in Go - 

```
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

```
var i int = 5   // declare and initialize

var f float32   // declare
f = i           // ERROR - implicit conversion not allowed in Go and will result in an error.
f = float32(i)  // Have to be explicit with the intent to convert to a different type.
```