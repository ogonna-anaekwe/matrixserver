package utils

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	const expected string = "1,2,3,4,5,6,7,8,9"
	given, _ := flatten(rows)
	if expected != given {
		t.Errorf("Expected %v. Got %v", expected, given)
	}
}
