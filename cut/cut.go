package cut

import (
	"errors"
	"strconv"
	"strings"
)

// Fields returns the requested field from the input, iterating over each line
func Fields(input string, fields string) (string, error) {
	// Variable to store output
	var (
		startStr, endStr string
		output           []string
	)

	// Convert fields to int
	requestedFieldParts := strings.Split(fields, "-")
	startStr = requestedFieldParts[0]
	endStr = startStr
	if len(requestedFieldParts) == 2 {
		endStr = requestedFieldParts[1]
	}

	start, err := strconv.Atoi(startStr)
	if err != nil {
		return "", errors.New("invalid field value")
	}
	end, err := strconv.Atoi(endStr)
	if err != nil {
		return "", errors.New("invalid field value")
	}

	if start < 1 || end < 1 {
		return "", errors.New("fields are numbered from 1")
	}

	// Split input into lines
	lines := strings.Split(input, "\n")

	// Iterate over lines
	for _, line := range lines {

		// Split line into fields
		lFields := strings.Fields(line)

		// Check if requested field is greater than the number of fields in the line
		if end > len(lFields) {
			continue
		}

		// Count of cut starts from 1, so we need to subtract 1 from given field
		startIndex := start - 1
		endIndex := end

		output = append(output, strings.Join(lFields[startIndex:endIndex], "\t"))
	}

	// Join output with new line and return
	return strings.Join(output, "\n"), nil
}
