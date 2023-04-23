package wctdd

import "strings"

func CountLines(input string) int {
	return strings.Count(input, "\n")
}

func CountWords(input string) int {
	return len(strings.Fields(string(input)))
}

func CountCharacters(input string) int {
	return strings.Count(string(input), "") - 1
}
