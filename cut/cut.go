package cut

import (
	"errors"
	"strconv"
	"strings"
)

// Fields returns the requested field from the input, iterating over each line
func Fields(input string, fields string) (string, error) {
	// Variable to store output
	var output []string

	// Convert fields to int
	fieldNum, err := strconv.Atoi(fields)
	if err != nil {
		return "", errors.New("invalid field value")
	}

	// Split input into lines
	lines := strings.Split(input, "\n")

	// Iterate over lines
	for _, line := range lines {

		// Split line into fields
		lFields := strings.Fields(line)

		// Check if requested field is greater than the number of fields in the line
		if fieldNum > len(lFields) {
			continue
		}

		// Count of cut starts from 1, so we need to subtract 1 from given field
		fieldIndex := fieldNum - 1

		// Append field to output
		output = append(output, lFields[fieldIndex])
	}

	// Join output with new line and return
	return strings.Join(output, "\n"), nil
}
