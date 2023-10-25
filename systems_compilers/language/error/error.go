package error

import (
	"fmt"
	"os"
)

var hadError bool = false

func CoolError(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[Line: %d] Error %s: %s\n", line, where, message)
	hadError = true
}

func ExitIfError() {
	if hadError {
		os.Exit(65)
	}
}
