package cut

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

const (
	// https://go.dev/doc/effective_go#printing
	MaxInt = int(^uint(0) >> 1) // largest int
)

// Fields returns the requested field from the input, iterating over each line
func Fields(input, fields, delimiter string) (string, error) {
	// Delimiter is tab by default
	if delimiter == "" {
		delimiter = "\t"
	}

	// Get the field ranges
	fRanges, err := fieldRanges(fields)
	if err != nil {
		return "", err
	}

	// Split input into lines
	lines := strings.Split(input, "\n")

	// Slice of string for storing each line's requested fields
	var finalOutput []string

	// Iterate over lines
	for _, line := range lines {
		// If line doesn't contain delimiter, add it to output without any changes
		if !strings.ContainsAny(line, delimiter) {
			finalOutput = append(finalOutput, line)
			continue
		}

		// Split line into fields
		lFields := strings.FieldsFunc(line, func(r rune) bool {
			return string(r) == delimiter
		})

		// Variable to store output for each line
		var (
			output []string
		)
		for _, r := range fRanges {
			// Count of cut starts from 1, so we need to subtract 1 from given field
			startIndex := r.start - 1
			endIndex := r.end

			// Check if requested field is greater than the number of fields in the line
			if startIndex > len(lFields) {
				continue
			}

			// If endIndex is MaxInt or greater than the number of fields present,
			// it means we want to cut till the end of the line
			if endIndex == MaxInt || endIndex > len(lFields) {
				endIndex = len(lFields)
			}

			// Just add the requested fields to output
			output = append(output, lFields[startIndex:endIndex]...)
		}

		// Join all the fields of a single line with the delimiter
		finalOutput = append(finalOutput, strings.Join(output, delimiter))
	}

	// For the final output, join each item in finalOutput with a newline
	return strings.Join(finalOutput, "\n"), nil
}

type fRange struct {
	start int
	end   int
}

func fieldRanges(f string) ([]fRange, error) {
	fieldRangesOutput := []fRange{}

	// Split all field ranges
	givenfieldRanges := strings.Split(f, ",")

	for _, fieldRange := range givenfieldRanges {
		var (
			startStr, endStr string
			start, end       int
		)

		// Split the range on `-`
		before, after, found := strings.Cut(fieldRange, "-")

		startStr = before
		endStr = after

		// If no `-`, that means it's not a range but a single field
		if !found {
			endStr = before
		}

		// If the field is passed as just `-`
		if startStr == "" && endStr == "" {
			return fieldRangesOutput, errors.New("invalid range with no endpoint")
		}

		// If the field is passed as `-X`, then consider the field as `1-X`
		if startStr == "" && endStr != "" {
			startStr = "1"
		}

		// If the field is passed as `X-`, then set end as MaxInt
		if startStr != "" && endStr == "" {
			endStr = strconv.Itoa(MaxInt)
		}

		// Convert to int
		start, err := strconv.Atoi(startStr)
		if err != nil {
			return fieldRangesOutput, errors.New("invalid field value")
		}

		// Convert to int
		end, err = strconv.Atoi(endStr)
		if err != nil {
			return fieldRangesOutput, errors.New("invalid field value")
		}

		// Field number can't be lower than 1
		if start < 1 || end < 1 {
			return fieldRangesOutput, errors.New("fields are numbered from 1")
		}

		fieldRangesOutput = append(fieldRangesOutput, fRange{
			start,
			end,
		})

	}

	// We want the output to be in order
	sort.Slice(fieldRangesOutput, func(i, j int) bool {
		return fieldRangesOutput[i].start < fieldRangesOutput[j].start
	})

	return fieldRangesOutput, nil
}
