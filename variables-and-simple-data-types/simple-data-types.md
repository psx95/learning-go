### Simple Data Types

There are only 4 simple data types 
 - Strings
 - Numbers
 - Booleans
 - Errors

 #### Strings

 Collection of 1 or more UTF-8 code points. There are 2 flavors - 
  - quoted string (delimited by double quotes) - represents an interpretted string
    ```
    "this is interpretted string, \n it interprets escape characters"

    If this is printed to the console, this will be the output - 

    this is interpretted string, 
    it interprets escape characters

    (Notice the newline)
    ```
  - backtick string (delimited by backticks) - represents a raw string
    ```
    `this is a raw string, \n it is represented by backquote marks`

    If this is printed to the console, this will be the output -

    this is a raw string, \n it is represented by backquote marks

    (Notice that the \n character is not interpreted and is printed as raw)

    Another feature of the raw string is, we can have carriage return in the middle of the string.
    For instance, the following string is completely valid in Go - 

    `raw strings
    ignore new line`

    This string if printed to console, will be displayed with the carriage return, generating - 

    raw strings
    ignore new lines
    ```

 #### Numbers

 There are 4 number types - 
 1. Integers
  - int : standard integers 
 2. Unsigned Integers
  - uint : these are whole numbers but cannot go negative
 3. Floating Point numbers
  - float32 : these are 32 bit floating point numbers
  - float64 : these are 64 bit floating point numbers
 4. Complex Numbers
    These contain a number representing a real part and an imaginary part
  - complex64 : these are comprised of 2 32-bit floating point numbers - a real and an imaginary
  - complex128 : these are comprised of 2 64-bit floating point numbers - a real and an imaginary
 
 #### Booleans

 Similar to booleans in other languages - there are 2 values `true` and `false` which are also first-class citizens in go which means they are not 
 represention of any numbers.

 #### Errors

The error type in go is an interface. It has the ability to report an error using a function named `Error()` which returns a string.

#### NOTES
1. The [builtin](https://pkg.go.dev/builtin) exposes a lot of the functionality that is built in to Go. This package does not need to be imported since its gloablly available in all
Go programs. This package also exposes all the simple data types that are mentioned above.