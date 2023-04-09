package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type fieldRange struct {
	start int
	end   int
}

type options struct {
	Delimiter string
	Fields    []fieldRange
	Output    io.Writer
}

func main() {
	// Define flags
	delimiterFlag := flag.String("delimiter", "\t", "delimiter for fields")
	fieldsFlag := flag.String("fields", "", "fields or ranges to cut (ex. 1,3-5)")

	flag.Parse()

	// If no flags are passed, print usage
	if flag.NFlag() == 0 {
		flag.Usage()
	}

	// Get operands
	files, err := getOperands(flag.Args())
	if err != nil {
		log.Fatal(err)
	}

	// Create options
	o := options{
		Delimiter: *delimiterFlag,
		Output:    os.Stdout,
	}

	// Convert fields string to field ranges which contain start and end as ints
	fRange, err := getFieldRanges(*fieldsFlag)
	if err != nil {
		log.Fatal(err)
	}

	o.Fields = fRange

	// Select fields from each file
	for _, file := range files {
		selectFields(file, o)
	}
}

// getFieldRanges returns a slice of field ranges after parsing the fields string.
func getFieldRanges(fields string) (options []fieldRange, err error) {
	for _, field := range strings.Split(fields, ",") {
		if strings.Contains(field, "-") {
			// If the field is a range, add it to the options
			fields := strings.Split(field, "-")
			start, err := strconv.Atoi(fields[0])
			if err != nil {
				return nil, err
			}
			end, err := strconv.Atoi(fields[1])
			if err != nil {
				return nil, err
			}
			options = append(options, fieldRange{
				start: start,
				end:   end,
			})
		} else {
			fieldInt, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			// If the field is not a range, add it to the options
			options = append(options, fieldRange{
				start: fieldInt,
				end:   fieldInt,
			})
		}
	}

	return options, nil
}

// selectFields filters the fields based on the field ranges and writes the selected fields to the output.
func selectFields(rd io.Reader, opts options) {
	// Selects the fields based on the field ranges
	// Writes the selected fields to the output
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		// Reads each line
		line := scanner.Text()

		// Splits the line string into fields based on the delimiter
		lineFields := strings.Split(line, opts.Delimiter)

		// Variable to hold the selected fields
		selections := []string{}

		// Loops through the given field ranges
		for _, field := range opts.Fields {

			// If the delimiter is not found in the line, return the entire line
			// This works because strings.Split() returns an array with the entire line
			// as the first element if the delimiter is not found
			if len(lineFields) == 1 {
				selections = append(selections, lineFields[0])
				break
			}

			if field.start == field.end {
				// If the field range is a single field, append it to the selections
				selections = append(selections, lineFields[field.start-1])
			} else {
				// If the field range is a range, append all fields in the range to the selections
				for i := field.start; i <= field.end; i++ {
					selections = append(selections, lineFields[i-1])
				}
			}
		}

		// Write the selections to the output appended with a newline
		opts.Output.Write([]byte(strings.Join(selections, opts.Delimiter) + "\n"))
	}
}

// getOperands returns a slice of file pointers after parsing the args.
func getOperands(args []string) (files []*os.File, err error) {
	if flag.NArg() == 0 {
		// If no arguments are passed, read from stdin.
		return []*os.File{os.Stdin}, nil
	} else if flag.NArg() == 1 && flag.Arg(0) == "-" {
		// If the argument is "-", read from stdin.
		return []*os.File{os.Stdin}, nil
	}

	// If arguments are passed, try opening them as files.
	if files, err = openFiles(flag.Args()); err != nil {
		return nil, err
	}

	return files, nil
}

// openFiles opens a list of files and returns a slice of file pointers.
func openFiles(filenames []string) ([]*os.File, error) {
	files := []*os.File{}
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}
