package main

import (
	"testing"
)

func TestFirstStar(t *testing.T) {
	var dict = map[string]int{
		">":          2,
		"^>v<":       4,
		"^v^v^v^v^v": 2,
	}
	for input, expected := range dict {
		var actual = firstStar(input)
		if actual != expected {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%s"`, actual, expected, input)
		}
	}
}

func TestSecondStar(t *testing.T) {
	var dict = map[string]int{
		"^v":         3,
		"^>v<":       3,
		"^v^v^v^v^v": 11,
	}
	for input, expected := range dict {
		var actual = secondStar(input)
		if actual != expected {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%s"`, actual, expected, input)
		}
	}
}
