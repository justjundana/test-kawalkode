# [![Go Report Card](https://goreportcard.com/badge/github.com/justjundana/jsonfmt)](https://goreportcard.com/report/github.com/justjundana/jsonfmt)
# ![Build](https://github.com/justjundana/jsonfmt/actions/workflows/test.yml/badge.svg)
# ![Go Version](https://img.shields.io/badge/go-1.21-blue)
# ![License](https://img.shields.io/badge/license-MIT-green)
# ![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
# jsonfmt

A modern CLI tool for formatting and validating JSON files with robust error handling and 100% test coverage.

## Features
- Pretty-print (format) JSON files
- Validate JSON files against JSON Schema
- Simple CLI: `jsonfmt <file.json>`
- Overwrites the file in-place
- Clear error messages on invalid or unreadable files
- Fully tested, 100% code coverage

## Installation

```sh
go install github.com/justjundana/jsonfmt/cmd/jsonfmt@latest
```

## Usage

```sh
jsonfmt <file.json>
```

### Current CLI Limitations
- Accepts **only one JSON file** as an argument (no multi-file support yet)
- No additional flags/options (`--help`, `--version`, `--output`, etc. are not available)
- The file is **overwritten in-place** (no output to stdout)
- On error (invalid file, read/write failure), an error message is printed to stderr and the program exits with a non-zero code

> Note: More features and options are planned for future releases.

## Example

**Before:**
```json
{"foo":1,"bar":[2,3]}
```

**After running:**
```sh
jsonfmt file.json
```

**After:**
```json
{
  "foo": 1,
  "bar": [
    2,
    3
  ]
}
```

## Testing

Run all tests (with 100% coverage):

```sh
go test -cover ./...
```

## Contributing

Contributions, issues, and feature requests are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Created by [@justjundana](https://github.com/justjundana) — feel free to reach out for questions or collaboration.
