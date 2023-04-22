package cut

import (
	"errors"
	"strconv"
	"strings"
)

// Fields returns the requested field from the input, iterating over each line
func Fields(input, fields, delimiter string) (string, error) {
	// Delimiter is tab by default
	if delimiter == "" {
		delimiter = "\t"
	}

	// Variable to store output
	var (
		startStr, endStr string
		start, end       int
		output           []string
	)

	// Convert fields to int
	requestedFieldParts := strings.Split(fields, "-")
	startStr = requestedFieldParts[0]
	endStr = startStr
	if len(requestedFieldParts) == 2 {
		endStr = requestedFieldParts[1]
	}

	if startStr == "" && endStr == "" {
		return "", errors.New("invalid range with no endpoint")
	}

	if startStr == "" && endStr != "" {
		startStr = "1"
	}

	start, err := strconv.Atoi(startStr)
	if err != nil {
		return "", errors.New("invalid field value")
	}

	if startStr != "" && endStr == "" {
		endStr = "inf"
		end = -1
	} else {
		end, err = strconv.Atoi(endStr)
		if err != nil {
			return "", errors.New("invalid field value")
		}
	}

	if start < 1 || (end < 1 && end != -1) {
		return "", errors.New("fields are numbered from 1")
	}

	// Split input into lines
	lines := strings.Split(input, "\n")

	// Iterate over lines
	for _, line := range lines {
		// If line doesn't contain delimiter, add it to output without any changes
		if !strings.ContainsAny(line, delimiter) {
			output = append(output, line)
			continue
		}

		// Split line into fields
		lFields := strings.FieldsFunc(line, func(r rune) bool {
			return string(r) == delimiter
		})

		// Count of cut starts from 1, so we need to subtract 1 from given field
		startIndex := start - 1
		endIndex := end

		// Check if requested field is greater than the number of fields in the line
		if startIndex > len(lFields) {
			output = append(output, "")
			continue
		}

		if endIndex > len(lFields) {
			output = append(output, "")
			continue
		}

		// If endIndex is -1, it means we want to cut till the end of the line
		if endIndex == -1 {
			endIndex = len(lFields)
		}

		output = append(output, strings.Join(lFields[startIndex:endIndex], delimiter))
	}

	// Join output with new line and return
	return strings.Join(output, "\n"), nil
}
