- use `go test` to run test.
- you can also specify the file name if not the `go test` command would try to run the test in the current directory.
- just shows you the time taken to run tests and `ok` if all tests pass.
- `ok      github.com/victor-nach/user-management-go/tests 0.883s`.
- if a test fails, the remaining tests would still run.

## -v flag
- `go test -v` means verbose output
- log all tests as they are run. Also print all text from Log and Logf calls even if the test succeeds.
- it shows you the tests as they are run.
- it also shows you the ones that passed and the ones that failed.
- if the test fails, it show's you the reason why `t.Errorf("Error message that shows in the terminal", got)`.
- The `-v` show's the name of the test function.
- Also show's if the test passed or not.
=== RUN   TestAbe
--- FAIL: TestAbe (0.00s)
    api_test.go:27: Error message that shows in the terminal
=== RUN   TestAr
--- PASS: TestAr (0.00s)

## cover
- Enable coverage analysis.
- adds a line to end of the code output containing the coverage `coverage: [no statements]`.

## failfast
- Do not start new tests after the first test failure.

## -coverprofile
- writes a coverage profile to a specific file.
- also sets the `-cover` flag.
- to use the flag you can either just pass the file name as the next parameter `-coverprofile cover.out`
- or use `-coverprofile=cover.out`

## go tool cover
- using the `go tool cover` command on the cli we can make use of the `.out` file created by cover
- we can use `go tool cover -html=c.out` to open a web browser displaying annotated source code
- `go tool cover -html=c.out -o coverage.html` to write out an HTML file instead of launching a web browser
- `go tool cover -func=c.out` to display coverage percentages to stdout for each function


