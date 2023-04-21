package cut

import (
	"strconv"
	"strings"
)

func Fields(input string, fields string) string {
	var output []string
	fieldNum, _ := strconv.Atoi(fields)
	fieldIndex := fieldNum - 1
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		lFields := strings.Fields(line)
		output = append(output, lFields[fieldIndex])
	}

	return strings.Join(output, "\n")
}
