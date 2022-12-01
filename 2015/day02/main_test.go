package main

import (
	"testing"
)

func TestPaperForBox(t *testing.T) {
	var dict = map[string]int{
		"2x3x4":  58,
		"1x1x10": 43,
	}
	for input, expected := range dict {
		var actual = paperForBox(input)
		if actual != expected {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%s"`, actual, expected, input)
		}
	}
}

func TestRibbonForBox(t *testing.T) {
	var dict = map[string]int{
		"2x3x4":   34,
		"1x1x10":  14,
		"9x23x15": 3153,
	}
	for input, expected := range dict {
		var actual = ribbonForBox(input)
		if actual != expected {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%s"`, actual, expected, input)
		}
	}
}
