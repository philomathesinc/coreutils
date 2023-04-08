package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	allArgs := os.Args[1:]
	fileName := allArgs[0]

	f, err := os.Open(fileName)
	if err != nil {
		erroredExit(err)
	}

	bytes := make([]byte, 1024)
	for {
		n, err := f.Read(bytes)
		if err == io.EOF {
			break
		}
		if err != nil {
			erroredExit(err)
		}
		fmt.Println(string(bytes[:n]))
	}
}

func cleanExit() {
	os.Exit(0)
}

func erroredExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
