package main

import (
	"flag"
	"fmt"
	"io"
	"log"
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

	filename := flag.Arg(0)
	if filename == "" || filename == "-" {
		tempFile, err := os.CreateTemp(".", "stdin-*")
		if err != nil {
			log.Fatalf("error creating temp file : %v", err)
		}

		file, err := os.Open("/dev/stdin")
		if err != nil {
			log.Fatalf("error opening stdin file : %v", err)
		}

		_, err = io.Copy(tempFile, file)
		if err != nil {
			log.Fatalf("error opening stdin file : %v", err)
		}

		tempFile.Close()
		file.Close()
		filename = tempFile.Name()
	}

	outputFormat := ""
	for i := 0; i < flag.NFlag(); i++ {
		if lineCountFlag {
			lineCount, err := CountLines(filename)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			outputFormat += fmt.Sprintf(" %d", lineCount)
			lineCountFlag = false
		}

		if wordCountFlag {
			wordCount, err := CountWords(filename)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			outputFormat += fmt.Sprintf(" %d", wordCount)
			wordCountFlag = false
		}

		if characterCountFlag {
			characterCount, err := CountCharacters(filename)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			outputFormat += fmt.Sprintf(" %d", characterCount)
			characterCountFlag = false
		}
	}
	fmt.Printf("%s %s\n", outputFormat, filename)
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
