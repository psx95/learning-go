### Maps in Go
Similar to an array, but allows for custom index (or *key*) types to reference values held in it.  Like slices, they are dynamically sized and are reference types. 

#### Important Notes
 - Since maps are reference types, they are pass by reference. So copying one map to another does not actually create a new map.
 It is pointing to the same underlying data structure to which the original map is pointing. 
 - To make deep copies of maps, use the `maps.Clone` method of the experimental maps package.
 - Go does not support performing comparisons on reference types. So comparing 2 maps via some comparison operator like `==` 
 will lead to a compile time error in Go. 
 - Ordering in a map is not determinant. 

#### Common Map Operations in Go
```go
// Declaring a map in Go - this will not initialize a map
var m map[string]int        // the value in the [] indicates the key type and "int" indicates value type
fmt.Println(m)              // map[] - like slices, this indicates a nil value

// Use the builtin make method to create an initialized map
initializedMap := make(map[string]int)

var referemcemap map[string][]string  // a map of string -> slice of strings

// Declaring and initializing a map 
m = map[string]int{"foo": 1, "bar" : 2}     // creates a map with 2 key-value entries 
fmt.Println(m)              // map[foo: 1, bar: 2]

// Looking up values from a Map
fmt.Println(m["foo"])       // returns 1

// Updating value in a Map
m["foo"] = 100
fmt.Println(m["foo"])       // returns 100

// Deleting values from a Map
delete(m, "foo")            // builtin function to delete entry with a given key

// Adding value to an existing map
m["baz"] = 98               // simply add a key that doesn't exist yet
fmt.Println(m)              // map[bar:2, baz:98] - only 2 since foo was deleted

// Accessing a key that does not exist in the map
fmt.Println(m["foo"])       // 0 - Go returns the zero value for the value type 

// But the above can be confusing since the map could have a key "foo" with a 
// value = 0. To avoid this, we can use the 'comma ok syntax'

v, ok := m["foo"]       // ok verifies the presence of the key in the map
fmt.Println(v, ok)      // 0, false - ok is false indicating that key "foo" wasn't found
// using ok as a variable name is just convention, technically, any variable name can be used
```
