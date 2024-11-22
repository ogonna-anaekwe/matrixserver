package utils

import (
	"net/http"
	"testing"
)

const testDataLocation string = "../../matrix.csv"

var (
	w       http.ResponseWriter
	rows, _ = ReadFile(testDataLocation, w)
)

func TestAllValuesDefined(t *testing.T) {
	_, err := allValuesDefined(rows)
	if err != nil {
		t.Errorf("Matrix dimension is missing value: %v", err)
	}
}

func TestIsSquareMatrix(t *testing.T) {
	_, err := isSquareMatrix(rows)
	if err != nil {
		t.Errorf("Matrix dimension is square: %v", err)
	}
}

func TestValidate(t *testing.T) {
	_, err := validate(rows)
	if err != nil {
		t.Errorf("Invalid matrix: %v", err)
	}
}

func TestReadFile(t *testing.T) {
	const expected int = 3
	given := len(rows)
	if expected != given {
		t.Errorf("Expected %v rows. Got %v rows", expected, given)
	}
}
