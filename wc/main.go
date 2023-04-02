package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]

	lineCount, err := CountLines(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(lineCount)
}

func CountLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	dataString := string(data)
	lines := strings.Split(dataString, "\n")

	return len(lines), nil
}
