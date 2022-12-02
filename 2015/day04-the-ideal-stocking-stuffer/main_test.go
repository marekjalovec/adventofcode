package main

import (
	"testing"
)

func TestGetSuffix(t *testing.T) {
	var dict = map[string]int{
		"abcdef":  609043,
		"pqrstuv": 1048970,
	}
	for input, expected := range dict {
		var actual = getSuffix(input, "00000")
		if actual != expected {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%s"`, actual, expected, input)
		}
	}
}
