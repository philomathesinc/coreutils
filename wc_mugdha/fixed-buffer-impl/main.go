package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	bufferSize = 1024
)

var (
	lineCount int
	charCount int
	wordCount int
	wordEnd   uint8
)

func main() {
	allArgs := os.Args[1:]

	// Going over each file
	for _, fileName := range allArgs {
		lineCount, wordCount, charCount = 0, 0, 0
		f := openFile(fileName)

		bytes := make([]byte, bufferSize)
		wordEnd = 0

		// Iterating over file contents progressively until EOF
		for {
			// Reading file contents into buffer
			n, err := f.Read(bytes)
			if err == io.EOF {
				break
			}
			if err != nil {
				erroredExit(err)
			}
			content := string(bytes[:n])

			// Counting lines, words and characters
			lineCount += strings.Count(content, "\n")
			wordCount += len(strings.Fields(content)) - int(wordEnd)
			charCount += len(content)

			// Checking if buffer ends with space or newline character
			if strings.HasSuffix(content, " ") || strings.HasSuffix(content, "\n") {
				wordEnd = 0
			} else {
				wordEnd = 1
			}

			// Debug statements - needs to be removed
			fmt.Printf("\nFields are: %q", strings.Fields(content))
			fmt.Printf("\nWordCount: %d", wordCount)
			fmt.Printf("\nLineCount: %d\n", lineCount)

		}

		f.Close()

		// Displaying the count for a specific file
		fmt.Println(lineCount, wordCount, charCount, fileName)
	}
}

func erroredExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func openFile(f string) *os.File {
	fd, err := os.Open(f)
	if err != nil {
		erroredExit(err)
	}
	return fd
}
