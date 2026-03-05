package internal

import (
	"fmt"
	"os"
)

// LogError prints an error message to stderr.
// Warning: format should be a constant string; user input should be passed via args to avoid format injection.
func LogError(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "[ERROR] %s\n", msg)
}

// LogInfo prints an info message to stdout.
// Warning: format should be a constant string; user input should be passed via args to avoid format injection.
func LogInfo(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stdout, "[INFO] %s\n", msg)
}
