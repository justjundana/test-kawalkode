package internal

import "fmt"

// ErrInvalidJSON is returned when JSON is invalid
var ErrInvalidJSON = fmt.Errorf("invalid JSON")

// ErrFileNotFound is returned when a file does not exist
var ErrFileNotFound = fmt.Errorf("file not found")

// ErrSchemaValidation is returned when schema validation fails
var ErrSchemaValidation = fmt.Errorf("schema validation failed")
