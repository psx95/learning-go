### Functions in Go
There are 3 main concepts about functions in Go - 
1. Function Signatures
2. Parameters & Arguments
3. Returning Values

#### Fucntion Signature
This is the code that describes the functions and its components. The function signature consists of the *function name, parameters and the return values*.

```go
// Here is what a function signature looks like - 

func functionName (functionParameters) (returnValues) {
    // function body
}
```

#### Passing parameters in functions in Go
Functions in Go can accept multiple parameters via a comma delimited list. The datatype of the function parameter comes after the variable within the function signature.
Various ways for passing parameters in Go have been discussed in the below examples - 

```go
func greet(name1 string, name2 string) {
    fmt.Println(name1)
    fmt.Println(name2)
}

// Note: If the function parameters listed adjecent to each other have the same type, then the datatype of the function parameters have to be defined only for the 
// last parameter. This means, the above function can also be written as - 

// Here name1 and name2 both have the string type
func greet(name1, name2 string) {
    fmt.Println(name1)
    fmt.Println(name2)
}

// Variadic Parameters - this is a special type of parameter that we can have AT THE END OF THE PARAMETER LIST of a function in Go.
// These are declared using the ellipsis operator (...) which converts the parameter from a single entity of a certain type to a collection of entities of that type.
// This is made clear using the below example -  
// names has now been coverted to a collection of strings and any number of strings can be passed to this function (comma separated).
func greet(names ...string) {
    // Inside this function, Go converts all the passed comma separated strings to a slice
    // However, the callers of this function cannot pass in slices - they need to pass the comma separated list of individual values
    for _, name := range names {
        fmt.Println(name)
    }
}
```
##### Difference between passing values and pointers to functions

This is the same concept as pass by value and pass by reference - when we pass values to a function, they are pass by value which means that any change the function makes on those
passed values is actually happening on a copy of the passed variable, and has no effect on the value being outside the function. 

Pointers on the other hand, represent pass by reference mechanism and the parameters passed this way are being shared between the called function and the calling function. This means 
that any change in the parameter made by the called function will take effect in the calling function as well.

To decide what mechanism to use - 
```
Use pointers only when you want to share memory, otherwise use values
```

#### Returning data from functions in Go

Functions in Go can return single, multiple or no values. The data types of the return values must be decalred in the function signature and the values from the functions must be
returned in an order that matches those data types. 

If multiple values are returned from a function, then the return statement must provide all the values. Omitting return values is not allowed.

```go
// Multiple return data types need to be enclosed in a parenthesis.
// This function returns an int  and a boolean value
func divide(l, r int) (int, bool) {
    if r == 0 {
        return 0, false // all values must be returned
    }
    // values are always returned in the order of function signature
    return l/r, true
}

// CONCEPT: NAMED RETURN TYPES
// Named return type allows us to name the return values recieved from the function - 
// This allows us to perform naked returns
func divide(l, r int) (result int, ok bool) {
    if r == 0 {
        // naked return
        // when go runs into this, it will return the current values for 'result' and 'ok' - which are 0 and false at this point
        return  // 0, false
    }
    // the name(s) of the variable here matches the name(s) given to the return value(s) in the function signature.
    result = l/r
    ok = true
    // this is called a naked return and is only possible when using named return types
    // however, even with named return types this is optional we could instead return the values similar to how 
    // they are done above even with named return types
    return // returns l/r, true
}

func main() {
    result, ok := divide(1, 2) // the variables are also pulled back in the order of function signature
    // result is implicitly assigned the int value from the function return 
    // ok is implicitly assigned the bool value from the function return
}
```
