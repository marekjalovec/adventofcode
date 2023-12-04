package main

import (
	"regexp"
	"testing"

	"adventofcode/utils"
)

func getLines() []string {

	lines := utils.ReadFile("input_test.txt")
	r1 := regexp.MustCompile(` +`)
	r2 := regexp.MustCompile(`^Card (\d+):`)
	for k, _ := range lines {
		lines[k] = r1.ReplaceAllString(lines[k], " ")
		lines[k] = r2.ReplaceAllString(lines[k], "")
	}
	return lines
}

func TestFirstStar(t *testing.T) {
	var expected = 13
	var actual = firstStar(getLines())

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}

func TestSecondStar(t *testing.T) {
	var expected = 30
	var actual = secondStar(getLines())

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}
