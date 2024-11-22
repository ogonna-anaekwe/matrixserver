package utils

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Valid operations.
const (
	Echo     string = "echo"
	Sum      string = "sum"
	Multiply string = "multiply"
	Flatten  string = "flatten"
	Invert   string = "invert"
)

// Validates incoming requests and dispatches to the relevant operation
func Reduce(rows [][]string, operation string) (string, error) {
	switch operation {
	case Echo:
		return echo(rows)
	case Sum:
		return sum(rows)
	case Multiply:
		return multiply(rows)
	case Flatten:
		return flatten(rows)
	case Invert:
		return invert(rows)
	default:
		return "-1", fmt.Errorf("unrecognized operation %v", operation)
	}
}

// Determines if the matrix is square.
func isSquareMatrix(rows [][]string) (bool, error) {
	numRows := len(rows)
	for i := 0; i < numRows; i++ {
		numCols := len(rows[i])
		if numCols != numRows {
			return false, fmt.Errorf("mismatch in matrix dimensions: %v cols and %v rows", numCols, numRows)
		}
	}

	return true, nil
}

// Determines if the matrix has non-empty cells.
func allValuesDefined(rows [][]string) (bool, error) {
	for rowIdx, row := range rows {
		for valIdx, val := range row {
			_, err := strconv.ParseInt(val, 10, 0)
			if err != nil {
				return false, fmt.Errorf("undefined value in matrix at row %v col %v", rowIdx, valIdx)
			}
		}
	}

	return true, nil
}

// Checks that the matrix is square and that every cell is populated.
func validate(rows [][]string) (bool, error) {
	_, err := isSquareMatrix(rows)
	if err != nil {
		return false, err
	}

	_, err = allValuesDefined(rows)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Reads file containing matrix data.
func ReadFile(location string, w http.ResponseWriter) ([][]string, error) {
	file, err := os.Open(location)
	if err != nil {
		logrus.Fatalf("Error opening file at %v %v\n", location, err)
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		res := fmt.Sprintf("error reading CSV file %s %s", location, err)
		w.Write([]byte(res))
		return nil, nil
	}

	_, err = validate(rows)
	if err != nil {
		return nil, fmt.Errorf("invalid input data %v:  %v", rows, err)
	}

	return rows, nil
}
