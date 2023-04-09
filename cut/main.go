package main

import (
	"flag"
	"log"
)

func main() {
	// Define flags
	delimiterFlag := flag.String("d", "\t", "delimiter for fields")
	fieldsFlag := flag.String("f", "", "fields or ranges to cut (ex. 1,3-5)")

	flag.Parse()

	// If no flags are passed, print usage
	if flag.NFlag() == 0 {
		flag.Usage()
	}

	log.Println("delimiterFlag:", *delimiterFlag, "fieldsFlag:", *fieldsFlag)
}
