### Arrays in Go
Arrays in Go function exactly same like arrays in any other language. The basic difference is in the syntax. 

#### Important Notes
 - Arrays start at zero index in Go
 - Assigning arrays in Go is a copy operation. In other words, they are pass-by-value.
 - Comparing two arrays of different data types or size will lead to a compile-time error in Go.

#### Common Array Operations in Go
```go
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

### Slices in Go
Slices in Go are homogenous data structures similar to an array, but they are dynamic in nature - they can grow or shrink in size. Also, unlike an array they do not hold their own data. They always point to data being held somewhere by an array &rarr; this also implies that they are reference types in Go. 

Since slices are backed by an internal array and reflects the values held in that array, any change in the underlying array will reflect a change in the slice and vice versa. 

#### Common Slice Operations in Go
```go
// Declaring Slices in Go
var slice []int             // Unlike an array, we do not declare a size

// Printing the above slice will result in [] 
fmt.Println(slice)          // [] (nil - a declared an uninitialized slice contains nil) 

// Declaring and initializing
var slice = []int{1, 2, 3}   // This will create a slice big enough to hold 3 values 

// Reading from a slice
fmt.Println(slice[1])   // 2 - Similar to an array

// Updating a value in slice
slice[1] = 0
fmt.Println(slice)  // [1 0 3]

// Adding values to slices
// first argument is the slice itself, then add as many elements as you wish
slice = append(slice, 6, 7, 8)      
fmt.Println(slice)  // [1 0 3 6 7 8]

// Deleting values from slices using experimental slices package
// first argument is the slice itself, the second argument is start index for delete (inclusive)
// and the third argument is ent end index for delete (exclusive)
slice = slices.Delete(s, 1, 3)
fmt.Println(slice)     // [1 6 7 8] - values 0 and 3 were deleted
```
