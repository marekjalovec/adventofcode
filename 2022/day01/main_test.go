package main

import "testing"

func TestFirstStar(t *testing.T) {
	var expected = 24000
	var actual = getSumForTopN("input_test.txt", 1)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}

func TestSecondStar(t *testing.T) {
	var expected = 45000
	var actual = getSumForTopN("input_test.txt", 3)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}
