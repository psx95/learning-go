### Contexts in Go
Contexts in Go were added as a means to convey the cancellations/completion of goroutines to other related goroutines effectively. So if a gourinte spaws multiple other goroutines which should all cancel/stop executing if the parent goroutine has stopped, this can be done using the context.
Communicating the cancellation effectively and early is possible so that unnecessary goroutines do not keep running and waste resources.

In Go, we derive contexts from existing contexts to form a context tree, so if an upstream context cancels, all the derived contexts also recieve that cancellation signal and are cancelled.

#### Creating contexts in Go
As mentioned above, in Go, we create contexts from existing contexts. The `context` package provide us three methods to create derived contexts -  `WithCancel`, `WithDeadline` and `WithTimeout`. All of these functions take in an existing context as the parameter.

Since the `context.Context` is an interface type in Go, in order to get existing context, Go provides us with two public methods -
1. `context.Background` - This represents the context that exists throughout the execution of the program. This means that this context never cancels or more accurately, when it does cancel, the program ends.
2. `context.TODO` - This is technically same as background but serves as a placeholder context. It serves as a documentation to the developers that a derived context should be used instead, but since that is not implemented/available right now, the background context will be used.

#### Caveats when using Contexts
1. Canceling contexts doesn't necessarily mean that the goroutines should immediately terminate what they're doing.
2. Canceling contexts doesn't cause the goroutines using that context to stop. Canceling contexts, merely sends a signal to all the child contexts and it is the responsibility of the programmer to actually check for the cancelation and exit the goroutine.

#### Conventions around the context in Go
When using contexts in Go, there are the following conventions around its use -
1. The context variable should be named `ctx`. If there are multiple contexts, the names should prefix with `ctx`. 
2. If a function/method accepts context as an argument, it should be the first argument that is declared in the method/function signature.
3. In a context aware program, contexts must always be passed as arguments to the goroutines.

#### Using contexts in Go
To see samples of how contexts are used in code checkout -
1. [context with cancel](../demo-programs/context-cancel-example/main.go) sameple.
