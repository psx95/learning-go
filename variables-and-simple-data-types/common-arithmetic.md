### Arithmetic & Comparisons Operators

#### Arithmetic Operators
Go supports all basic arithmetic operators - 

```go
a, b := 10, 5       // a and b both are inferred here to be of type int

// Addition
c := a + b          // c = 15 is also of the type int
// Subtraction
c = a - b           // c = 5 
// Multiplication
c = a * b           // c = 59
// Division
c = a / b           // c = 2 
c = a / 3           // c = 3 : Go supports integer division, but remainder will be lost here
d := 7.0 / 2.0      // d = 3.5 : Go supports floating point division as well
// Modulus
c = a % 3           // c = 1
```
The above are just the basic types. An exhaustive list can be found here - [Language Specification - Arithmetic Operators](https://go.dev/ref/spec#Arithmetic_operators)

#### Comparison Operators
Go also has support for basic comparison operators - 

1. Equality operator (`==`)
    Go supports equality comparison for all value types except errors (they do not check for equality the same way as for other types)
2. Inequality operator (`!=`)
     Go supports inequality comparison for all value types except errors (they do not check for equality the same way as for other types)
3. Less than (`<`)
4. Less than or equal to (`<=`)
5. Greater than (`>`)
6. Greater than or equal to (`>=`)

Note: Operators 3 to 6 are only supported for numeric types. 

For more details on comparisons, check [Language Specification - Comparison Operators](https://go.dev/ref/spec#Comparison_operators)