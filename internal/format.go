package internal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

// MinifyJSONMarshalOnly marshals a user-supplied struct (for test coverage of marshal error branch)
func MinifyJSONMarshalOnly(v interface{}) ([]byte, error) {
	minified, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("cannot minify JSON: %w", err)
	}
	return minified, nil
}

// MinifyJSONFileReturnTyped unmarshals into a user-supplied struct, then marshals it (for full coverage and advanced use cases)
func MinifyJSONFileReturnTyped(path string, v interface{}) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return MinifyJSONMarshalOnly(v)
}

// ValidateJSONWithSchema memvalidasi file JSON terhadap file schema JSON
func ValidateJSONWithSchema(jsonPath, schemaPath string) error {
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
	documentLoader := gojsonschema.NewReferenceLoader("file://" + jsonPath)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("invalid schema: %w", err)
	}
	if !result.Valid() {
		msg := ""
		for _, desc := range result.Errors() {
			msg += desc.String() + "; "
		}
		return fmt.Errorf("schema validation failed: %s", msg)
	}
	return nil
}

// MarshalIndentFunc mendefinisikan signature fungsi marshal indentasi
type MarshalIndentFunc func(v interface{}, prefix, indent string) ([]byte, error)

// FormatJSONFileWith memformat file JSON dengan fungsi marshal custom
func FormatJSONFileWith(path string, marshal MarshalIndentFunc) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	pretty, err := marshal(v, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("cannot format JSON: %w", err)
	}
	if err := os.WriteFile(path, pretty, 0644); err != nil {
		return nil, fmt.Errorf("cannot write file: %w", err)
	}
	return pretty, nil
}

// FormatJSONFile memformat file JSON dengan indentasi
func FormatJSONFile(path string) error {
	_, err := FormatJSONFileWith(path, json.MarshalIndent)
	return err
}

// FormatJSONFileWithReturn memformat file JSON dan mengembalikan hasil []byte (tanpa menulis file)
func FormatJSONFileWithReturn(path string, marshal MarshalIndentFunc) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	pretty, err := marshal(v, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("cannot format JSON: %w", err)
	}
	return pretty, nil
}

// MinifyJSONFile menulis file JSON hasil minify
func MinifyJSONFile(path string) error {
	out, err := MinifyJSONFileReturn(path)
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, out, 0644); err != nil {
		return fmt.Errorf("cannot write file: %w", err)
	}
	return nil
}

// MinifyJSONFileReturn mengembalikan hasil minify []byte (tanpa menulis file)
func MinifyJSONFileReturn(path string) ([]byte, error) {
	return MinifyJSONFileReturnWith(path, json.Marshal)
}

// MinifyJSONFileReturnWith allows injecting a custom marshal function for testing
func MinifyJSONFileReturnWith(path string, marshal func(v interface{}) ([]byte, error)) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	minified, err := marshal(v)
	if err != nil {
		return nil, fmt.Errorf("cannot minify JSON: %w", err)
	}
	return minified, nil
}
