package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/philomathesinc/coreutils/cut"
)

func main() {
	// Define flags
	delimiterFlag := flag.String("d", "\t", "delimiter for fields")
	fieldsFlag := flag.String("f", "", "fields or ranges to cut (ex. 1,3-5)")

	flag.Parse()

	// If no flags are passed, print usage
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	files := flag.Args()

	if flag.NArg() == 0 {
		files = append(files, "-")
	}

	for _, filename := range files {
		var (
			file *os.File
			err  error
		)

		if filename == "-" {
			file = os.Stdin
		} else {
			file, err = os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: No such file or directory", filename)
			}
			defer file.Close()
		}

		bytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read file")
		}

		fields, err := cut.Fields(string(bytes), *fieldsFlag, *delimiterFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		fmt.Println(fields)

	}
}
