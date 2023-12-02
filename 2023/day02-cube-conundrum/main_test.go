package main

import (
	"testing"

	"adventofcode/utils"
)

var lines = utils.ReadFile("input_test.txt")

func TestFirstStar(t *testing.T) {
	var expected = 8
	var actual = firstStar(lines)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}

func TestSecondStar(t *testing.T) {
	var expected = 2286
	var actual = secondStar(lines)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}
