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
	fieldIndex := fieldNum - 1
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		lFields := strings.Fields(line)
		output = append(output, lFields[fieldIndex])
	}

	return strings.Join(output, "\n"), nil
}
