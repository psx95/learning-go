### Packages in Go
 1. Packages provide a higher level mechanism for organize the Go applications or Go modules. 
 2. Package is simply a directory within a module - since it can be the root directory, the module can itself serve as a package. 
 3. Packages in Go must contain at least one go source file.
 4. All source files within a package can share data with each other. In other words, all members within a package are visible to each other.

*NOTE: As a reminder, a Go module is any directory with `go.mod` file inside it.*

#### Important environment variables in Go
 - GOROOT: Location where Go is installed on your system. This is used to find Go binaries and libraries.
 - GOPATH: Location of the Go workspace. This is where all the Go code resides.

#### Using packages in Go

##### Package declaration
The very first line in all go source files declare the package. `package main` conveys a special meaning to the Go compiler, but all the other package names need to be within a folder with the same name. The name of the folder and the package name declaration must match - this is used by the Go compiler to determine where to look for the go source files.
So, for instance - consider the following go source file
```go
// somefile.go
package user

func foo() {
    // function code
}
```
**This file must be present in a directory called user, otherwise the program will not compile.**

As mentioned above, `package main` conveys a special meaning to the compiler, so any source go files using this package, need not be in a directory named `main`. 


##### Using other packages
In order to use code/functionality defined in other packages in Go, we use the `import packagename`.

##### Package level members and visibility

In a Go source file, package level members are any members that are declared outside any function scope in the source file. For instance,

```go
// somefile.go
package user

// MaxUsers constant is a package level members
const MaxUsers = 20

// foo is a package level member
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

Go supports two levels of visibility - 
1. Package level visibility (the lowest level of visibility available)
    - They are defined by naming the identifier starting with a lowercase letter. Eg. `user`
    - Such fields are reserved for use within the package.
    - Any package level member can see any other package level member.    
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

##### Recommended folder structure for GOROOT and GOPATH
 - GOROOT: should contain the following folders
  - src: contains source files for the standard library
  - pkg: contains compiled packages used by Go
  - bin: contains executable commands

- GOPATH: should contain the following folders
 - src: all Go source files
 - pkg: contains third party libraries or binaries
 - bin: contains all the executables

#### Initializing packages - the init function
When running a Go program, the Go runtime looks for the main function and executes it directly. However, before the Go runtime executes the main function, it runs the `init()` function automatically.

The `init()` function is responsible for performing initialization tasks such as setting up logging or creating global variables or any other tasks that must be done before the program starts. The Init function has the following properties - 
1. Called by Go runtime
2. Takes no argument
3. No return value
4. Runs before the main function
5. There can be multiple init functions in a Go application - they are run according based on the file names in alphabetical order
6. Each init function is executed only once per application at the start of the program

```go
// main.go
import "fmt"
func init() {
    fmt.Println("Hello from init")
}

func main() {
    fmt.Println("Hello from main")
}
// When this file is run, the following output is expected to be seen:
// Hello from init
// Hello from main
```
