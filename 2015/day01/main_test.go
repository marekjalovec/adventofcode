package main

import "testing"

func TestFirstStar(t *testing.T) {
	var dict = map[string]int{
		"(())":    0,
		"()()":    0,
		"(((":     3,
		"(()(()(": 3,
		"))(((((": 3,
		"())":     -1,
		"))(":     -1,
		")())())": -3,
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
		")":     1,
		"()())": 5,
		"(":     -1,
	}
	for input, expected := range dict {
		var actual = secondStar(input)
		if actual != expected {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%s"`, actual, expected, input)
		}
	}
}
