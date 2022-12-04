package main

import (
	"adventofcode/utils"
	"testing"
)

var lines = utils.ReadFile("input_test.txt")

func TestFirstStar(t *testing.T) {
	var expected = 2
	var actual = firstStar(lines)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}

func TestSecondStar(t *testing.T) {
	var expected = 4
	var actual = secondStar(lines)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}
