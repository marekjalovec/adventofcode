package main

import (
	"testing"

	"adventofcode/utils"
)

var lines1 = utils.ReadFile("input_test1.txt")
var lines2 = utils.ReadFile("input_test2.txt")

func TestFirstStar(t *testing.T) {
	var expected = 142
	var actual = firstStar(lines1)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}

func TestSecondStar(t *testing.T) {
	var expected = 281
	var actual = secondStar(lines2)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}
