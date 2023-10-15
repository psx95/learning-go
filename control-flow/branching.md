### Branching in Go
Go provides us with the following mechanisms to introduce execution flow branching in our programs - 
1. if statements
2. switch statements
3. deferred functions
4. panic and recovery
5. goto statements

#### If statements
```go
// simple if statement
// "test" is a statement or expression that evaluates to a boolean value
// if "test" evaluates to true, then the set of statements is executed, otherwise not
if test {
    // set of statements
}

// if-else statement
// if "test" evaluates to true, then the if block is executed, otherwise else block is executed.
// else block is optional
if test {
    // set of statements
} else {
    // set of statements
}

// if-else-if ladder - can have any number of else-if conditions 
// only one of the branches gets executed here - depending on whichever test evaluates to true
// the test conditions are evaluated in order, so if a branch's test evaluates to true, the branches 
// defined below it will be skipped
if test {
    // set of statements
} else if test2 {
    // set of statements
} else if test3 {
    // set of statements
} else {
    // set of statements
}

// if statement with optional initializer
// Just like the for loop, the if statement also allows for an optional initializer
// this initializer statement executes once before evaluating the test condition
// the initializer is only allowed on the opening if statement - i.e, it is not allowed 
// on the else-if statements
if initializer; test {
    // set of statements
}
```

#### Switch statements
*Specialized form of `if` statements.*

```go
// The test_expression here should evaluate to a single value and then that value is matched
// up with the expressions defined in the cases and set of statements defined in that case block
// are executed
switch test_expression {
    case expression1:
        // set of statements
    case expression2, expression3:
        // set of statements
        // combining two expression values into a single case means that the statements that need to
        // be executed for all the values are same\
    default:
        // set of statements
        // default case is executed in case no expression value matches the test_expression
}

// Switch statement also allows for an optional initializer - which executes before the test_expression is evaluated
switch initializer; test_expression {
    case expression1:
        // set of statements
    case expression2:
        // set of statements
    default:
        // set of statements
}

// EXAMPLE
switch i := 5; i {
    case 1:
        fmt.Println("first case")
    case 2 + 3, 2*i+3:
        // this will get executed since 2 + 3 evaluates to 5
        fmt.Println("second case")
    default:
        fmt.Println("default case")
}

// Logical Switches - test_expression is a boolean value
// Logical switches are syntactically no different from the simple switch statements shown above
// The difference comes from how they are used and a few syntactical shorthand
// Since test_expression is a boolean here, the cases have to evaluate to boolean values themselves
switch i := 8; true {
    case i > 5:
        fmt.Println("i is greater than 5")
    case i > 10:
        fmt.Println("i is greater than 10")
    default:
        fmt.Println("i is less than or equal to 5")
}

// Since logical expressions are commonly used in Go, there is an option to leave out the `true` from the test_expression
// In that case, the test expression will be an implied true and since the initializer is also optional, the above switch
// case could also be written as -
i := 8
switch { // this evaluates to "switch true"
    case i > 5:
        fmt.Println("i is greater than 5")
    case i > 10:
        fmt.Println("i is greater than 10")
    default:
        fmt.Println("i is less than or equal to 5")
}
```
