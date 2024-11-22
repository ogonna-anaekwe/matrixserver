package utils

import (
	"strconv"
	"testing"
)

func TestSum(t *testing.T) {
	const expected int64 = 45
	given, _ := sum(rows)
	_given, _ := strconv.ParseInt(given, 10, 0)
	if expected != _given {
		t.Errorf("Expected %v. Got %v", expected, _given)
	}
}
