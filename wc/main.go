package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	lineCountFlag bool
	wordCountFlag bool
)

func main() {
	flag.BoolVar(&lineCountFlag, "l", false, "Display the number of lines")
	flag.BoolVar(&wordCountFlag, "w", false, "Display the number of words")
	flag.Parse()
	filename := os.Args[2]

	if lineCountFlag {
		lineCount, err := CountLines(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(lineCount, filename)
	}

	if wordCountFlag {
		lineCount, err := CountWords(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(lineCount, filename)
	}
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

	return strings.Count(string(data), "\n"), nil
}

func CountWords(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	return strings.Count(string(data), " "), nil
}
