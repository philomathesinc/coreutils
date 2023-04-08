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
	output    string
)

func main() {
	allArgs := os.Args[1:]

	for _, fileName := range allArgs {
		f, err := os.Open(fileName)
		if err != nil {
			erroredExit(err)
		}
		defer f.Close()

		bytes := make([]byte, 1024)
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
			charCount += len(content)
			wordCount += len(strings.Split(content, " ")) + lineCount
		}

		output = fileName

		fmt.Println(output)
	}
}

func cleanExit() {
	os.Exit(0)
}

func erroredExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
