package utils

import (
	"strings"
)

// Unnests matrix.
func flatten(rows [][]string) (string, error) {
	var res strings.Builder
	for rowIdx, row := range rows {
		for valIdx, val := range row {
			res.WriteString(val)
			addComma(rows, row, rowIdx, valIdx, &res)
		}
	}

	return res.String(), nil
}

// Adds trailing comma to cells.
func addComma(rows [][]string, row []string, rowIdx int, valIdx int, res *strings.Builder) {
	isLastValue := ((rowIdx + 1) * (valIdx + 1)) == (len(rows) * len(row))
	if !isLastValue {
		res.WriteString(",")
	}
}
