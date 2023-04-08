package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	lineCountFlag      bool
	wordCountFlag      bool
	characterCountFlag bool
)

func main() {
	flag.BoolVar(&lineCountFlag, "l", false, "Display the number of lines")
	flag.BoolVar(&wordCountFlag, "w", false, "Display the number of words")
	flag.BoolVar(&characterCountFlag, "c", false, "Display the number of characters")
	flag.Parse()

	if !lineCountFlag && !wordCountFlag && !characterCountFlag {
		filename := os.Args[1]
		lineCount, err := CountLines(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		wordCount, err := CountWords(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		characterCount, err := CountCharacters(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%d %d %d %s\n", lineCount, wordCount, characterCount, filename)
		return
	}

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
		wordCount, err := CountWords(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(wordCount, filename)
	}

	if characterCountFlag {
		characterCount, err := CountCharacters(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(characterCount, filename)
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

	return len(strings.Fields(string(data))), nil
}

func CountCharacters(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	return strings.Count(string(data), "") - 1, nil
}
