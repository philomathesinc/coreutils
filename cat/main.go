package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.Arg(0)

	files, err := getFiles(flag.Args())

	fmt.Printf("%s %s\n", , filename)

}
