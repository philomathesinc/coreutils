package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()
	output := Cat(flag.Args())
	fmt.Println(output)
}

func Cat(filenames []string) string {
	var outputdata []byte
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		data := readData(file)
		outputdata = append(outputdata, data...)
	}
	return string(append(outputdata, '%'))
}

func readData(r io.Reader) []byte {
	data, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
