package utils

import (
	"testing"
)

func TestEcho(t *testing.T) {
	expected := "1,2,3\n4,5,6\n7,8,9\n"
	given, _ := echo(rows)
	for i := range given {
		if expected[i] != given[i] {
			t.Errorf("Expected %v. Got %v", string(expected[i]), string(given[i]))
		}
	}
}
