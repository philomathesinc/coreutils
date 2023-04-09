package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	lineCount int
	charCount int
	wordCount int
	wordEnd   uint8
)

func main() {
	allArgs := os.Args[1:]

	for _, fileName := range allArgs {
		lineCount, wordCount, charCount = 0, 0, 0
		f, err := os.Open(fileName)
		if err != nil {
			erroredExit(err)
		}
		defer f.Close()

		bytes := make([]byte, 64)
		wordEnd = 0

		for {
			n, err := f.Read(bytes)
			if err == io.EOF {
				break
			}
			if err != nil {
				erroredExit(err)
			}
			content := string(bytes[:n])

			lineCount += strings.Count(content, "\n")
			wordCount += len(strings.Fields(content)) - int(wordEnd)
			charCount += len(content)

			if strings.HasSuffix(content, " ") || strings.HasSuffix(content, "\n") {
				wordEnd = 0
			} else {
				wordEnd = 1
			}

			fmt.Printf("\nFields are: %q", strings.Fields(content))
			fmt.Printf("\nWordEnd: %d", wordEnd)
			fmt.Printf("\nWordCount: %d", wordCount)
			fmt.Printf("\nLineCount: %d\n", lineCount)
		}

		fmt.Println(lineCount, wordCount, charCount, fileName)
	}
}

func cleanExit() {
	os.Exit(0)
}

func erroredExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
