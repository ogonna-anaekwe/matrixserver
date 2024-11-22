package utils

import (
	"strconv"
	"testing"
)

func TestMultiply(t *testing.T) {
	const expected int64 = 362880
	given, _ := multiply(rows)
	_given, _ := strconv.ParseInt(given, 10, 0)
	if expected != _given {
		t.Errorf("Expected %v. Got %v", expected, _given)
	}
}
