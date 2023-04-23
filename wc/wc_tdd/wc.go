package wctdd

import "strings"

func CountLines(input string) int {
	return strings.Count(input, "\n")
}
