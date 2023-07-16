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
// Just like the for loop, the if statement also allows for an optional intitializer
// this initializer statement executes once before evaluating the test condition
// the initialzer is only allowed on the opening if statement - i.e, it is not allowed 
// on the else-if statements
if initializer; test {
    // set of statements
}
```
