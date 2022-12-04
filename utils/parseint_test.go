package utils

import "testing"

func TestParseInteger(t *testing.T) {
	var dict = map[string]int{
		"-10":  -10,
		"10": 10,
		"1024902384209423": 1024902384209423,
		"0": 0,
		"": 0,
		"foo": 0,
	}
	for input, expected := range dict {
		var actual = ParseInt(input)
		if actual != expected {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%s"`, actual, expected, input)
		}
	}
}
