### Packages in Go
 1. Packages provide a higher level mechansim for organize the Go applications or Go modules. 
 2. Package is simply a directory within a module - since it can be the root directory, the module can itself serve as a packge. 
 3. Packages in Go must contain at least one go source file.
 4. All source files within a package can share data with each other. In other words, all members within a package are visible to each other.

*NOTE: As a reminder, a Go module is any directory with `go.mod` file inside it.*

#### Using packages in Go

##### Package declaration
The very first line in all go source files declare the package. `package main` conveys a special meaning to the Go compiler, but all the other packge names need to be within a folder with the same name. The name of the folder and the package name declaration must match - this is used by the Go compiler to determine where to look for the go source files.
So, for instance - consider the following go source file
```go
// somefile.go
package user

func foo() {
    // function code
}
```
This file must be present in a directory called user, otherwise the program will not compile.
As mentioned above, `package main` conveys a special meaning to the compiler, so any source go files using this package, need not be in a directory named `main`. 


##### Using other packages
In order to use code/functionality defined in other packages in Go, we use the `import packagename`.

##### Package level members and visibility

In a Go source file, package level members are any  memebers that are declared outside any function scope in the source file. For instance,

```go
// somefile.go
package user

// MaxUsers constant is a package level memebrs
const MaxUsers = 20

// foo is a package level memeber
func foo() {
    // function code
}

// bar is a package level member
func bar() {
    // variable s is not a package level member
    var s string
    fmt.Println(s)
}
``` 

Package level members can be accessed by source Go files in other packages that import the package containing these members, provided the members have sufficient visibility.

Go supports two levels of visibiilty - 
1. Package level visibility (the lowest level of visibility available)
    - They are defined by naming the indentifier starting with a lowercase letter. Eg. `user`
    - Such fields are reserved for use within the package.
    - Any package level member can see any other package level memeber.    
2. Public level visibility
    - They are defined by naming the identifier starting with an uppercase letter. Eg. `User`
    - Such fields can be used by any package member - within the same package as well as outside of the package.
    - Basically, public fields can be used by anything that imports the package.

```go
// somefile.go
package user

// This is a public struct - so that means it can also be accessed outside user package
type User struct {
    ID       int        // public field: can be accessed outside the user package
    Username string     // public field: can be accessed outside the user package
    password string     // private member: reserved for use within the user package
}
```
