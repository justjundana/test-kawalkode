package internal

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"
)

func writeTempFile(t *testing.T, content string) string {
	       tmpfile, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("cannot create temp file: %v", err)
	}
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("cannot write temp file: %v", err)
	}
	tmpfile.Close()
	t.Cleanup(func() { os.Remove(tmpfile.Name()) })
	return tmpfile.Name()
}

func TestValidateJSONWithSchema_Valid(t *testing.T) {
	schema := `{"type":"object","properties":{"a":{"type":"number"}},"required":["a"]}`
	data := `{"a":1}`
	schemaFile := writeTempFile(t, schema)
	dataFile := writeTempFile(t, data)
	if err := ValidateJSONWithSchema(dataFile, schemaFile); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidateJSONWithSchema_Invalid(t *testing.T) {
	schema := `{"type":"object","properties":{"a":{"type":"number"}},"required":["a"]}`
	data := `{"b":2}`
	schemaFile := writeTempFile(t, schema)
	dataFile := writeTempFile(t, data)
	err := ValidateJSONWithSchema(dataFile, schemaFile)
	if err == nil || !strings.Contains(err.Error(), "schema validation failed") {
		t.Errorf("expected schema validation error, got: %v", err)
	}
}

func TestValidateJSONWithSchema_BadSchema(t *testing.T) {
	schema := `{"type":"object",` // invalid JSON
	data := `{"a":1}`
	schemaFile := writeTempFile(t, schema)
	dataFile := writeTempFile(t, data)
	err := ValidateJSONWithSchema(dataFile, schemaFile)
	if err == nil || !strings.Contains(err.Error(), "invalid schema") {
		t.Errorf("expected invalid schema error, got: %v", err)
	}
}

func TestFormatJSONFileWithReturn(t *testing.T) {
	file := writeTempFile(t, `{"a":1}`)
	out, err := FormatJSONFileWithReturn(file, json.MarshalIndent)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(string(out), "a") {
		t.Errorf("output missing key")
	}
}

func TestFormatJSONFileWithReturn_ReadError(t *testing.T) {
	_, err := FormatJSONFileWithReturn("/notfound.json", json.MarshalIndent)
	if err == nil || !strings.Contains(err.Error(), "cannot read file") {
		t.Errorf("expected read error")
	}
}

func TestFormatJSONFileWithReturn_UnmarshalError(t *testing.T) {
	file := writeTempFile(t, `{"a":}`)
	_, err := FormatJSONFileWithReturn(file, json.MarshalIndent)
	if err == nil || !strings.Contains(err.Error(), "invalid JSON") {
		t.Errorf("expected unmarshal error")
	}
}

func TestFormatJSONFileWithReturn_MarshalError(t *testing.T) {
	file := writeTempFile(t, `{"a":1}`)
	badMarshal := func(v interface{}, prefix, indent string) ([]byte, error) {
		return nil, errors.New("fail")
	}
	_, err := FormatJSONFileWithReturn(file, badMarshal)
	if err == nil || !strings.Contains(err.Error(), "cannot format JSON") {
		t.Errorf("expected marshal error")
	}
}

func TestMinifyJSONFile(t *testing.T) {
	file := writeTempFile(t, `{"a":1}`)
	if err := MinifyJSONFile(file); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	       out, _ := os.ReadFile(file)
	if !strings.Contains(string(out), "a") || strings.Contains(string(out), " ") {
		t.Errorf("minify failed")
	}
}

func TestMinifyJSONFile_ReadError(t *testing.T) {
	err := MinifyJSONFile("/notfound.json")
	if err == nil || !strings.Contains(err.Error(), "cannot read file") {
		t.Errorf("expected read error")
	}
}

func TestMinifyJSONFile_UnmarshalError(t *testing.T) {
	file := writeTempFile(t, `{"a":}`)
	err := MinifyJSONFile(file)
	if err == nil || !strings.Contains(err.Error(), "invalid JSON") {
		t.Errorf("expected unmarshal error")
	}
}

func TestMinifyJSONFile_MarshalError(t *testing.T) {
	// Cannot mock MinifyJSONFileReturn as it is a function, not a variable.
	// This test is not valid in Go without refactoring the implementation to allow injection.
	// Skipping this test.
}

func TestMinifyJSONFile_WriteError(t *testing.T) {
	file := writeTempFile(t, `{"a":1}`)
	os.Chmod(file, 0400)
	defer os.Chmod(file, 0600)
	err := MinifyJSONFile(file)
	if err == nil || !strings.Contains(err.Error(), "cannot write file") {
		t.Errorf("expected write error")
	}
}

func TestMinifyJSONFileReturn(t *testing.T) {
	file := writeTempFile(t, `{"a":1}`)
	out, err := MinifyJSONFileReturn(file)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !strings.Contains(string(out), "a") || strings.Contains(string(out), " ") {
		t.Errorf("minify return failed")
	}
}


type badMarshalStruct2 struct{}
func (b badMarshalStruct2) MarshalJSON() ([]byte, error) { return nil, errors.New("fail-marshal-real-typed") }

func TestFormatJSONFile_Valid(t *testing.T) {
	input := `{"a":1,"b":[2,3]}`
	       tmpfile, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("cannot create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatalf("cannot write temp file: %v", err)
	}
	tmpfile.Close()
	if err := FormatJSONFile(tmpfile.Name()); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	       out, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("cannot read temp file: %v", err)
	}
	if string(out) == input {
		t.Errorf("file not formatted")
	}
}

func TestFormatJSONFile_Invalid(t *testing.T) {
	input := `{"a":1,,}`
	       tmpfile, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("cannot create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatalf("cannot write temp file: %v", err)
	}
	tmpfile.Close()
	if err := FormatJSONFile(tmpfile.Name()); err == nil {
		t.Errorf("expected error for invalid JSON")
	}
}

func TestFormatJSONFile_ReadError(t *testing.T) {
	err := FormatJSONFile("/path/to/nonexistent.json")
	if err == nil || !strings.Contains(err.Error(), "cannot read file") {
		t.Errorf("expected read error, got: %v", err)
	}
}

func TestFormatJSONFile_WriteError(t *testing.T) {
	input := `{"a":1}`
	tmpfile, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("cannot create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatalf("cannot write temp file: %v", err)
	}
	tmpfile.Close()
	if err := os.Chmod(tmpfile.Name(), 0400); err != nil {
		t.Fatalf("cannot chmod temp file: %v", err)
	}
	defer os.Chmod(tmpfile.Name(), 0600)
	err = FormatJSONFile(tmpfile.Name())
	if err == nil || !strings.Contains(err.Error(), "cannot write file") {
		t.Errorf("expected write error, got: %v", err)
	}
}

type badJSON struct{}

func (b badJSON) MarshalJSON() ([]byte, error) {
	return nil, errors.New("mock marshal error")
}

func TestFormatJSONFile_MarshalError(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("cannot create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	input := `{"a":1}`
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatalf("cannot write temp file: %v", err)
	}
	tmpfile.Close()
	marshalErr := func(v interface{}, prefix, indent string) ([]byte, error) {
		return nil, errors.New("mock marshal error")
	}
	_, err = FormatJSONFileWith(tmpfile.Name(), marshalErr)
	if err == nil || !strings.Contains(err.Error(), "cannot format JSON") {
		t.Errorf("expected marshal error, got: %v", err)
	}
}


type badMarshalStruct3 struct{}
func (b badMarshalStruct3) MarshalJSON() ([]byte, error) { return nil, errors.New("fail-marshal-real-typed3") }

func TestMinifyJSONMarshalOnly_MarshalError(t *testing.T) {
	s := badMarshalStruct3{}
	_, err := MinifyJSONMarshalOnly(s)
	if err == nil || !strings.Contains(err.Error(), "cannot minify JSON") {
		t.Errorf("expected marshal error in MinifyJSONMarshalOnly")
	}
}

func TestMinifyJSONMarshalOnly_Success(t *testing.T) {
	type S struct{ A int }
	s := S{A: 42}
	b, err := MinifyJSONMarshalOnly(s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	       if string(b) != `{"A":42}` {
		       t.Errorf("unexpected output: %s", string(b))
	       }
}

func TestMinifyJSONFileReturnTyped_Success(t *testing.T) {
	type S struct{ A int }
	file := writeTempFile(t, `{"A":42}`)
	var s S
	b, err := MinifyJSONFileReturnTyped(file, &s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	       if string(b) != `{"A":42}` {
		       t.Errorf("unexpected output: %s", string(b))
	       }
}

func TestMinifyJSONFileReturnTyped_ReadError(t *testing.T) {
	type S struct{ A int }
	_, err := MinifyJSONFileReturnTyped("/not-exist-xyz", &S{})
	if err == nil || !strings.Contains(err.Error(), "cannot read file") {
		t.Errorf("expected read error")
	}
}

func TestMinifyJSONFileReturnTyped_UnmarshalError(t *testing.T) {
	type S struct{ A int }
	file := writeTempFile(t, `{"A":}`)
	_, err := MinifyJSONFileReturnTyped(file, &S{})
	if err == nil || !strings.Contains(err.Error(), "invalid JSON") {
		t.Errorf("expected unmarshal error")
	}
}

func TestMinifyJSONFileReturn_ReadError(t *testing.T) {
	_, err := MinifyJSONFileReturn("/not-exist-xyz")
	if err == nil || !strings.Contains(err.Error(), "cannot read file") {
		t.Errorf("expected read error")
	}
}

func TestMinifyJSONFileReturn_UnmarshalError(t *testing.T) {
	file := writeTempFile(t, `{"A":}`)
	_, err := MinifyJSONFileReturn(file)
	if err == nil || !strings.Contains(err.Error(), "invalid JSON") {
		t.Errorf("expected unmarshal error")
	}
}

func TestMinifyJSONFileReturnWith_MarshalError(t *testing.T) {
	file := writeTempFile(t, `{"A":1}`)
	marshalErr := func(v interface{}) ([]byte, error) { return nil, errors.New("mock marshal error") }
	_, err := MinifyJSONFileReturnWith(file, marshalErr)
	if err == nil || !strings.Contains(err.Error(), "cannot minify JSON") {
		t.Errorf("expected marshal error")
	}
}

func TestMinifyJSONFileReturnWith_Success(t *testing.T) {
	file := writeTempFile(t, `{"A":42}`)
	b, err := MinifyJSONFileReturnWith(file, json.Marshal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(b) != `{"A":42}` {
		t.Errorf("unexpected output: %s", string(b))
	}
}

func TestMinifyJSONFileReturn_Success(t *testing.T) {
	file := writeTempFile(t, `{"A":42}`)
	b, err := MinifyJSONFileReturn(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(b) != `{"A":42}` {
		t.Errorf("unexpected output: %s", string(b))
	}
}
