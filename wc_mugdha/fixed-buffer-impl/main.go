package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	lineCountFlag      bool
	wordCountFlag      bool
	characterCountFlag bool
)

func main() {
	var allFiles []string
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-l" || os.Args[i] == "-w" || os.Args[i] == "-c" {
			continue
		}
		allFiles = append(allFiles, os.Args[i])
		os.Args = append(os.Args[:i], os.Args[i+1:]...)
		i--
	}

	flag.BoolVar(&lineCountFlag, "l", false, "Display the number of lines")
	flag.BoolVar(&wordCountFlag, "w", false, "Display the number of words")
	flag.BoolVar(&characterCountFlag, "c", false, "Display the number of characters")
	flag.Parse()

	var printTotal bool
	lc := make(map[string]string)
	wc := make(map[string]string)
	cc := make(map[string]string)
	var tl, tw, tc = "", "", ""

	if !lineCountFlag && !wordCountFlag && !characterCountFlag {
		lineCountFlag, wordCountFlag, characterCountFlag = true, true, true
	}

	if len(allFiles) > 1 {
		printTotal = true
	}

	if lineCountFlag {
		total := 0
		for _, fileName := range allFiles {
			lcInt := countLines(fileName)
			lc[fileName] = strconv.Itoa(lcInt)
			total += lcInt
		}
		tl = fmt.Sprintf("%s ", strconv.Itoa(total))
	}
	if wordCountFlag {
		total := 0
		for _, fileName := range allFiles {
			wcInt := (countWords(fileName))
			wc[fileName] = strconv.Itoa(wcInt)
			total += wcInt
		}
		tw = fmt.Sprintf("%s ", strconv.Itoa(total))
	}
	if characterCountFlag {
		total := 0
		for _, fileName := range allFiles {
			ccInt := countChars(fileName)
			cc[fileName] = strconv.Itoa(ccInt)
			total += ccInt
		}
		tc = fmt.Sprintf("%s ", strconv.Itoa(total))
	}

	for _, k := range allFiles {
		fmt.Printf("%*s%*s%*s %s\n", len(tl), lc[k], len(tw), wc[k], len(tc), cc[k], k)
	}

	if printTotal {
		fmt.Printf(" %s%s%s%s", tl, tw, tc, "total\n")
	}
}

func countLines(fn string) int {
	f, err := os.Open(fn)
	if err != nil {
		erroredExit(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var c int
	sc.Split(bufio.ScanBytes)
	for sc.Scan() {
		if sc.Text() == "\n" {
			c++
		}
	}
	if err != nil {
		erroredExit(err)
	}

	return c
}

func countWords(fn string) int {
	f, err := os.Open(fn)
	if err != nil {
		erroredExit(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var c int
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		c++
	}
	if err != nil {
		erroredExit(err)
	}

	return c
}

func countChars(fn string) int {
	f, err := os.Open(fn)
	if err != nil {
		erroredExit(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var c int
	sc.Split(bufio.ScanBytes)
	for sc.Scan() {
		c++
	}
	if err != nil {
		erroredExit(err)
	}

	return c
}

func erroredExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func openFile(f string) *os.File {
	fd, err := os.Open(f)
	if err != nil {
		erroredExit(err)
	}
	return fd
}
