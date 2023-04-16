package main

import (
	"flag"
	"log"
	"os"

	"github.com/philomathesinc/coreutils/wc"
)

func main() {
	var (
		lineCountFlag      bool
		wordCountFlag      bool
		characterCountFlag bool
		filenames          []string
	)
	flag.BoolVar(&lineCountFlag, "l", false, "Display the number of lines")
	flag.BoolVar(&wordCountFlag, "w", false, "Display the number of words")
	flag.BoolVar(&characterCountFlag, "c", false, "Display the number of characters")
	flag.Parse()

	filenames = flag.Args()

	file, err := os.Open(filenames[0])
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	_ = wc.New(lineCountFlag, characterCountFlag, wordCountFlag, file)
}
