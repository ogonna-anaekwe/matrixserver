package utils

import (
	"fmt"
	"strconv"

	l "github.com/ogonna-anaekwe/matrixserver/internal/logger"
)

// Computes product of all values in the matrix.
func multiply(rows [][]string) (string, error) {
	log := l.NewLogger()
	log.SetReportCaller(true)

	var res int64 = 1
	for _, row := range rows {
		for _, val := range row {
			v, err := strconv.ParseInt(val, 10, 0)
			if err != nil {
				log.Errorf("error computing product for %v %v \n", rows, err)
			}
			res *= v
		}
	}

	return fmt.Sprintf("%d", res), nil
}
