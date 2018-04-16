package logs

import "fmt"

var (
	testname = "Network"
)

// StartTest starts scope of a test.
func StartTest(name string) {
	testname = name
}

// Fail prints message in failure format.
func Fail(message string) {
	fmt.Printf("[FAILED] [%s] %s\n", testname, message)
}

// Succeed prints the message in passed format.
func Succeed(message string) {
	fmt.Printf("[PASSED] [%s] %s\n", testname, message)
}

// Info prints the message in info format.
func Info(message string) {
	fmt.Printf("[INFO] [%s] %s\n", testname, message)
}
