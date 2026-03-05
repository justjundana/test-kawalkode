package internal

import (
	"fmt"
	"os"
)

// LogError prints an error message to stderr.
// Accepts only pre-formatted messages to prevent format string injection.
func LogError(msg string) {
	fmt.Fprintf(os.Stderr, "[ERROR] %s\n", msg)
}

// LogInfo prints an info message to stdout.
// Accepts only pre-formatted messages to prevent format string injection.
func LogInfo(msg string) {
	fmt.Fprintf(os.Stdout, "[INFO] %s\n", msg)
}
