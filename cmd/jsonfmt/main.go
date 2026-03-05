package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"io/ioutil"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/justjundana/jsonfmt/internal"
)

func main() {
	minify := flag.Bool("minify", false, "Minify JSON (hilangkan whitespace)")
	schema := flag.String("schema", "", "Path ke file JSON Schema untuk validasi")
	toStdout := flag.Bool("stdout", false, "Output ke stdout, file tidak diubah")
	diffFlag := flag.Bool("diff", false, "Tampilkan diff berwarna antara file sebelum dan sesudah format/minify")
	help := flag.Bool("help", false, "Tampilkan bantuan")
	version := flag.Bool("version", false, "Tampilkan versi jsonfmt")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: jsonfmt [--minify] [--schema <schema.json>] [--stdout] [--diff] <file.json> [file2.json ...]\n")
	}
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if *version {
		fmt.Println("jsonfmt version 1.0.0")
		os.Exit(0)
	}
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(2)
	}
	exitCode := 0
	for i := 0; i < flag.NArg(); i++ {
		file := flag.Arg(i)
		var err error
		if *schema != "" {
			err = internal.ValidateJSONWithSchema(file, *schema)
		} else if *diffFlag {
			// diff mode: tampilkan diff berwarna, file tidak diubah
			orig, readErr := ioutil.ReadFile(file)
			if readErr != nil {
				fmt.Fprintf(os.Stderr, "Error on %s: %v\n", file, readErr)
				exitCode = 1
				continue
			}
			var out []byte
			if *minify {
				out, err = internal.MinifyJSONFileReturn(file)
			} else {
				out, err = internal.FormatJSONFileWithReturn(file, internal.MarshalIndentFunc(json.MarshalIndent))
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error on %s: %v\n", file, err)
				exitCode = 1
				continue
			}
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(string(orig), string(out), false)
			fmt.Print(dmp.DiffPrettyText(diffs))
		} else if *minify {
			if *toStdout {
				out, err2 := internal.MinifyJSONFileReturn(file)
				if err2 != nil {
					err = err2
				} else {
					os.Stdout.Write(out)
					os.Stdout.Write([]byte("\n"))
				}
			} else {
				err = internal.MinifyJSONFile(file)
			}
		} else {
			if *toStdout {
				out, err2 := internal.FormatJSONFileWithReturn(file, internal.MarshalIndentFunc(json.MarshalIndent))
				if err2 != nil {
					err = err2
				} else {
					os.Stdout.Write(out)
					os.Stdout.Write([]byte("\n"))
				}
			} else {
				err = internal.FormatJSONFile(file)
			}
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on %s: %v\n", file, err)
			exitCode = 1
		}
	}
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}
