package utils

import (
	"fmt"
	"strings"
)

// Prints contents of matrix.
func echo(rows [][]string) (string, error) {
	var res string
	for _, row := range rows {
		res = fmt.Sprintf("%s%s\n", res, strings.Join(row, ","))
	}

	return res, nil
}
