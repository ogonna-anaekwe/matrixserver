package utils

import (
	"testing"
)

func TestInvert(t *testing.T) {
	const expected string = "1,4,7\n2,5,8\n3,6,9\n"
	given, _ := invert(rows)
	for i := range given {
		if expected[i] != given[i] {
			t.Errorf("Expected %v. Got %v", string(expected[i]), string(given[i]))
		}
	}
}
