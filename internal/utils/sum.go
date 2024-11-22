package utils

import (
	"fmt"
	"strconv"

	l "github.com/ogonna-anaekwe/matrixserver/internal/logger"
)

// Computes sum of all values in the matrix.
func sum(rows [][]string) (string, error) {
	log := l.NewLogger()
	log.SetReportCaller(true)

	var res int64
	for _, row := range rows {
		for _, val := range row {
			v, err := strconv.ParseInt(val, 10, 0)
			if err != nil {
				log.Errorf("error computing sum for %v %v \n", rows, err)
			}
			res += v
		}
	}

	return fmt.Sprintf("%d", res), nil
}
