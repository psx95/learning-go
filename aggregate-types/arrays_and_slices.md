### Arrays in Go

Arrays in Go function exactly same like arrays in any other language. The basic difference is in the syntax. 

#### Important Notes
 - Arrays start at zero index in Go
 - Assigning arrays in Go is a copy operation. In other words, they are pass-by-value.
 - Comparing two arrays of different data types or size will lead to a compile-time error in Go.

#### Common Array Operations in Go
```
// Declaring Arrays in Go
var arr [3]int           // array of 3 ints

// Printing the above array will result in 0,0,0 
fmt.Println(arr)        // [0 0 0]

// Declaring and initializing arrays
arr = [3]int{1, 2, 3}

// Reading from array in Go
fmt.Println(arr[1])     // 2

// Updating values in an array
arr[1] = 5
fmt.Println(arr)        // [1 5 3]

// Getting length of an array in Go
 fmt.Println(len(arr))   // 3

// Comparing two arrays
arr2 = [3]int{1, 2, 3}

// Go compares the 2 arrays sequentially, 1 element at a time.
arr == arr2             // false - since arr[1] = 5
```