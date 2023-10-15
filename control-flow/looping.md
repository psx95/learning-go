### Looping in Go
All loops in Go are `for` loops, but there are different forms of the loop available - 
 - infinite loops
 - loop till condition
 - counter based loops

 #### Important Notes
  - `break` statement can be used to exit early from the loops.

 #### Common loop operations in Go
 ```go
 // infinite loops - there is no external reason as to why this loop may stop
 for  {
    // set of statements
 }
 // EXAMPLE - INFINITE LOOP
 i := 1
 for { 
    // This loop will never exit - the program may crash once it runs out of valid integers
    // Or the system can run into OOM issues - external factors may affect the running of loops
    // break statement can be used to exit from this loop
    fmt.Println(i)
    i += 1      // short-hand form for update and re-assign
 }

 // loop till condition - loop continues till the condition remains true
 // "condition" can be a Go statement or expression that evaluates to a boolean
 for condition {
  // set of statements  
 }
 // EXAMPLE - LOOP TILL CONDITION
 j := 1
 for j < 3 { // Loop exits as soon as j becomes 3
    fmt.Println(j)
    j += 1
 } 
 fmt.Println("DONE")

 // counter based loops - 
 // "initializer" is a statement that is executed before the start of the loop 
 // "condition" is same as defined in the loop till condition - if condition evaluates, program does not enter the loop
 // "post clase" this statement is executed after going through the loop body - typically used to increment a counter variable
 for initializer; condition; post clause {
    // set of statements
 }
 // EXAMPLE - COUNTER BASED LOOPS
 for k := 1; k < 3; k++ {
    fmt.Println(k)
 }
 fmt.Println("Done")
 ```

 #### Looping through collections
 Unlike basic looping where we can loop till some condition evaluates to false, when looping through collections, Go has built-in signals that
 indicate when the program has iterated over every member of the collection once.

 In Go, the `range` keyword indicates that we are looping through collections. The collections here can be - 
 1. arrays
 2. slices
 3. maps

 *There is no built-in way in Go to loop through structs.*

 ```go
 // various forms of looping over collections
 // looping over the collection - retrieving both key and value for each member of the collection
 // key for arrays and slices is the index number
 for key, value := range collection { ... } 

 // looping over the collection - retrieving only the key for each member of the collection
 for key := range collection { ... }
 // looping over the collection - retrieving only the value for each member of the collection
 // the underscore is called the blank identifier and signals Go that the program identifies that key is coming, but is not used by the program
 for _, value := range collection { ... } 

 arr := int[3]{1, 2, 3}
 for index, value := range arr {
    fmt.Println(index, value)
    // The loop will automatically terminate once it has gone through every value in the array
 }
 fmt.Println("Done")
 ```