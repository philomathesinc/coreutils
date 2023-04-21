package cut

import (
	"errors"
	"strconv"
	"strings"
)

func Fields(input string, fields string) (string, error) {
	var output []string
	fieldNum, err := strconv.Atoi(fields)
	if err != nil {
		return "", errors.New("invalid field value")
	}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		lFields := strings.Fields(line)
		if fieldNum > len(lFields) {
			continue
		}
		fieldIndex := fieldNum - 1
		output = append(output, lFields[fieldIndex])
	}

	return strings.Join(output, "\n"), nil
}
