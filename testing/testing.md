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
    - To fail a test we can report an error on the paramater `t *testing.T` like - `t.Errorf()` 
    - If there is no failure, test is assumed to have passed.
6. To run Go tests -
    ```shell
        # To run tests in the current directory
        go test .

        # To run all tests in module from the root directory of module
        # This will run all tests in the root directory as well as sub-directories.
        go test ./...    
    ```