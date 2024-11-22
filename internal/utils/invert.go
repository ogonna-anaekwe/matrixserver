package utils

import (
	"fmt"
	"strings"
)

// Transposes/Inverts a matrix.
func invert(rows [][]string) (string, error) {
	var res string
	numRows := len(rows)
	for rowIdx := range rows {
		var rowSlice []string
		for i := 0; i < numRows; i++ {
			rowSlice = append(rowSlice, rows[i][rowIdx])
		}
		res = fmt.Sprintf("%s%s\n", res, strings.Join(rowSlice, ","))
	}

	return res, nil
}
