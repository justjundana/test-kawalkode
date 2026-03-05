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
jsonfmt [--minify] [--schema <schema.json>] [--stdout] [--diff] <file.json>
```

### Flags

- `--minify`   Minify JSON (remove whitespace)
- `--schema`   Validate JSON file against a JSON Schema
- `--stdout`   Output to stdout instead of overwriting file
- `--diff`     Show color diff between original and formatted/minified
- `--help`     Show help message
- `--version`  Show version info

### Examples

Format a file:
```sh
jsonfmt data.json
```

Minify a file:
```sh
jsonfmt --minify data.json
```

Validate with schema:
```sh
jsonfmt --schema schema.json data.json
```

Show diff only:
```sh
jsonfmt --diff data.json
```

Output to stdout:
```sh
jsonfmt --stdout data.json
```

Show help:
```sh
jsonfmt --help
```

Show version:
```sh
jsonfmt --version
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
