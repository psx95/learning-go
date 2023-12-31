### Testing in Go
Go has built-in support for various kinds of testing for applications. 

#### To test correctness of business logic
1. Test (Unit test)
    - A regular unit test that tests the smallest part of an application such as a function in isolation.
2. Fuzz test
    - Uses AI to generate a plethora of inputs to provide to an otherwise normal test. This makes it easy to find edge case scenarios which might be missed when writing test cases.
3. Example test
    - Similar to a regular test, but has tight integration to Go documentation. This is used for examples used in technical documentation. These tests ensure that such examples are correct and then the Go documentation generator can put such examples in technical documentation.

##### To measure performance of the application code -
1. Benchmark Test 
    - Such tests allow us to setup a test scenario and then lets us confirm how much memory and time are required to execute that scenario in a given environment.

##### To measure resource usage of the application code - 
1. Profiling 
 - Takes a test scenario and gives us aggregate result on memory and cpu usage. Breaking it down further, by function showing where exactly that time and memory was spent.
2. Tracing
 - Similar to profiling, but also includes time to get a sense of how these resources were used in our application over time.

#### Rules for writing simple tests in Go
1. The name of the file containing tests should end with `_test.go`. This tells the Go compiler where to look for tests and exclude them when building application binary.
2. Typically the test file lives in the same package as the source file it is testing.
3. The test function in the test file starts with a prefix `Test`. It can be followed with any text but the text should start with a capital letter.
4. The test function receive a pointer to a testing object - `*testing.T`. This pointer is used to convey the status for the current test back to the Go runtime.
5. Unlike most other languages, Go does not have built-in assertions for the tests we write. 
    - To fail a test we can report an error on the paramter `t *testing.T` like - `t.Errorf()` 
    - If there is no failure, test is assumed to have passed.
6. To run Go tests -
    ```shell
        # To run tests in the current directory
        go test .
        # To run tests using fully qualified package ID (this uses module name)
        go test {modulename}/{directory_name}

        # To run all tests in module from the root directory of module
        # This will run all tests in the root directory as well as sub-directories.
        go test ./...

        # To run test functions that match a particular regex, use the run flag
        go test ./... run <regex>
        # Making a concrete example of the above, if we wanted to run a test function
        # whose name contains the word 'Add', we would issue the following command -
        go test ./... run Add

        # To skip selected tests from the test suite using the testing API (the t object)
        # Inside the test function, invoke the skip method - 
        t.Skip()

    ```

#### Ways of reporting test failures in Go
Go provides multiple ways to report a test failure, but they are broadly divided in two categories - immediate failures and non-immediate failures - 

##### Non-immediate failures
This means that the some test conditions failed, but the test is still valid and so other conditions may still be checked - the test function keeps running.
Methods to indicate this kind of failure -
1. `t.Fail()` - Simply marks test as failed and moves on
2. `t.Error(...interface{})` - Same as fail, but allows the user to pass additional information through the variadic interfaces.
3. `t.Errorf(string, ...interface{{})` - Same as error, but allows the user to pass a formatted string instead of raw input types.
*Note: Each of these methods have a corresponding method which results in the immediate failure of the test. This is discussed below.*

##### Immediate failures
This means that the test failed catastrophically and the testing function should not be allowed to continue executing.
Methods to indicate this kind of failure - 
1. `t.FailNow()` - corresponds to `t.Fail()`
2. `t.Fatal(...interface{})` - corresponds to `t.Error(...interface{})`
3. `t.Fatalf(string, ...interface{})` - corresponds to `t.Errorf(string, ...interface{})`

#### Black box testing vs White box testing
Black box testing in Go means writing tests using only the public API available. In contrast, when writing white box tests, we have full access to all the code we are testing.

These differences in the tests are merely conceptual knowledge. The Go test runner does not make distinction between the two kinds of tests - it will run both kinds of tests in the exact same manner.

All differences between the two kinds of testing are explained in the table below - 
| White Box Tests | Block Box Tests |
| --------------- | --------------- |
| Test has access to all the code | Test can only interact with the public API |
| Tests are in the same package as the production code | Tests are in a separate package |

*NOTE: For black box testing, we need the code to be in a separate package, which would typically mean creating another directory. However, Go recognizes this use case of writing black box tests and allows putting the black box test in a separate directory without creating a corresponding folder.*
