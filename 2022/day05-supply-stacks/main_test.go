package main

import (
	"testing"

	"adventofcode/utils"
)

var lines = utils.ReadFile("input_test.txt")

func getTestStacks() [][]string {
	return [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
}

func TestFirstStar(t *testing.T) {
	var expected = "CMZ"
	var actual = firstStar(lines, getTestStacks())

	if actual != expected {
		t.Fatalf(`Actual output %s does not match the expected output %s`, actual, expected)
	}
}

func TestSecondStar(t *testing.T) {
	var expected = "MCD"
	var actual = secondStar(lines, getTestStacks())

	if actual != expected {
		t.Fatalf(`Actual output %s does not match the expected output %s`, actual, expected)
	}
}
