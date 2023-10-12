### Race detector in Go
Go build tool has a built-in race detector that helps detects data races in concurrent Go programs.
It can be passed as a command line flag to the `go run` command - `-race` flag. This tells Go to instrument the build and look for data races.

An important thing to note here is that the race detector only detects races that occur in the program. This means that if there is a race condition that occurs only under certain circumstances or environment, which is not similar to the one during the build process, the race detector will not detect it.

Invoking the race detector is simple - 

```shell
# run from the directory containing the go source files
go run -race .
```
*Typically the `-race` flag is not used in production environment since it dramatically slows down the build process.*

In the [mutex-example](../demo-programs/mutex-example/main.go), a race condition can be created by removing the use of mutex locks.

Running the above sample without the mutex locks and the race flag should detect a data race condition. 