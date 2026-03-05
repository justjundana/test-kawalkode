package internal

import (
	"fmt"
	"os"
)

// LogError prints an error message to stderr
func LogError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "[ERROR] "+format+"\n", args...)
}

// LogInfo prints an info message to stdout
func LogInfo(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}
