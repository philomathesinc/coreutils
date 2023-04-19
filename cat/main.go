package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("cat")
	output := Cat(flag.Args())
	fmt.Println(output)
}

func Cat(filenames []string) (output string) {
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal("failed to open file:", err)
		}
		data, err := io.ReadAll(file)
		if err != nil {
			log.Fatal("failed to read file:", err)
		}
		file.Close()
		output += string(data)
	}
	return output
}
