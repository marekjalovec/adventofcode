package main

import (
	"adventofcode/utils"
	"testing"
)

var lines = utils.ReadFile("input_test.txt")

func TestFirstStar(t *testing.T) {
	var expected = 24000
	var actual = getSumForTopN(lines, 1)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}

func TestSecondStar(t *testing.T) {
	var expected = 45000
	var actual = getSumForTopN(lines, 3)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}
