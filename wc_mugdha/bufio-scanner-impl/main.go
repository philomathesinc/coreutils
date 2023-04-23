package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	lineCountFlag      bool
	wordCountFlag      bool
	characterCountFlag bool
)

func main() {
	var allFiles = flag.Args()
	// Parse over all CLI args to remove filenames from them before parsing flags.
	for i := 1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "-l") || strings.HasPrefix(os.Args[i], "-w") || strings.HasPrefix(os.Args[i], "-c") {
			continue
		}
		allFiles = append(allFiles, os.Args[i])
		os.Args = append(os.Args[:i], os.Args[i+1:]...)
		i--
	}

	// Separating the flags before parsing, if needed.
	var finalArgs []string
	for i := 1; i < len(os.Args); i++ {
		allFlags := strings.Split(os.Args[i], "")
		for _, v := range allFlags {
			if strings.EqualFold(v, "-") {
				continue
			}
			finalArgs = append(finalArgs, fmt.Sprintf("-%s", v))
		}
	}
	os.Args = append(os.Args[:1], finalArgs...)

	// Flag parsing.
	flag.BoolVar(&lineCountFlag, "l", false, "Display the number of lines")
	flag.BoolVar(&wordCountFlag, "w", false, "Display the number of words")
	flag.BoolVar(&characterCountFlag, "c", false, "Display the number of characters")
	flag.Parse()

	var printTotal bool
	lc := make(map[string]string)
	wc := make(map[string]string)
	cc := make(map[string]string)
	var tl, tw, tc = "", "", ""

	// Set all flags to true, if none are provided.
	if !lineCountFlag && !wordCountFlag && !characterCountFlag {
		lineCountFlag, wordCountFlag, characterCountFlag = true, true, true
	}

	// Printing total only if there are more than one files.
	if len(allFiles) > 1 {
		printTotal = true
	}

	// Counting lines if flag set.
	if lineCountFlag {
		total := 0
		for _, fileName := range allFiles {
			lcInt := countLines(fileName)
			lc[fileName] = strconv.Itoa(lcInt)
			total += lcInt
		}
		tl = fmt.Sprintf("%s ", strconv.Itoa(total))
	}

	// Counting words if flag set.
	if wordCountFlag {
		total := 0
		for _, fileName := range allFiles {
			wcInt := (countWords(fileName))
			wc[fileName] = strconv.Itoa(wcInt)
			total += wcInt
		}
		tw = fmt.Sprintf("%s ", strconv.Itoa(total))
	}

	// Counting characters if flag set.
	if characterCountFlag {
		total := 0
		for _, fileName := range allFiles {
			ccInt := countChars(fileName)
			cc[fileName] = strconv.Itoa(ccInt)
			total += ccInt
		}
		tc = fmt.Sprintf("%s ", strconv.Itoa(total))
	}

	// Printing count for provided files.
	for _, k := range allFiles {
		fmt.Printf("%*s%*s%*s %s\n", len(tl), lc[k], len(tw), wc[k], len(tc), cc[k], k)
	}

	// Printing total count.
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
