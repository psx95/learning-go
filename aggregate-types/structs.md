### Structs in Go
Struct is a fixed size aggregate type (similar to an array in this aspect) but its values can be of different types - making this a heterogenous data type.
Struct is the only heterogenous aggregate type available in Go.

#### Important Notes
 - Struct can contain any number of fields. Fields are elements within the struct. 
 - Each field has a name, a type and a value of that type. Fields can be of different types.
 - Structs have to be defined at compile time. Once defined, the struct cannot change at runtime.
 - Structs can contain other structs as their fields.
 - Structs are often used via custom types in Go to avoid the cumbersome syntax of declaring structs.
 - Structs are value types (they are pass by value) just like arrays.
 - Structs are comparable. 

#### Common Struct Operations in Go
```go
// Declaring an anonymous struct 
var s struct {
    name string
    id int
}
// Uninitialized struct (zero-value struct) has all its fields initialized to their zero values
fmt.Println(s)      // {"" 0}

// Updating fields of a struct
s.name = "psx95"
// Querying a struct field
fmt.Println(s.name) // psx95

// Creating custom type based on Struct
type myStruct struct {
    name string
    id int
}
// Declaring variable with custom type
var s myStruct
fmt.Println(s)      // {"" 0}
// Using struct literal to initialize a struct
s = myStruct {
    name: "psx"
    id: 95
}
fmt.Println(s)      // {"psx" 95}
s2 := s             // This will create a copy
s.name = "abc"
fmt.Prinltn(s, s2)  // {"abc" 95} {"psx" 95}

s == s2             // false - Go will check for same type and then check value of each field.
```