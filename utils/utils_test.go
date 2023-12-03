package utils

import (
	"testing"
)

func TestSum(t *testing.T) {
	var dict = []struct {
		values []int
		result int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{-5, 5}, 0},
		{[]int{}, 0},
	}
	for _, input := range dict {
		var actual = Sum(input.values)
		if actual != input.result {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%v"`, actual, input.result, input.values)
		}
	}
}

func TestParseInteger(t *testing.T) {
	var dict = map[string]int{
		"-10":              -10,
		"10":               10,
		"1024902384209423": 1024902384209423,
		"0":                0,
		"":                 0,
		"foo":              0,
	}
	for input, expected := range dict {
		var actual = ParseInt(input)
		if actual != expected {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%s"`, actual, expected, input)
		}
	}
}
