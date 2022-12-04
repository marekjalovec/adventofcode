package utils

import "testing"

type TestCase struct {
	values []int
	result int
}

func TestSum(t *testing.T) {
	var dict = []TestCase{
		{[]int{1, 2, 3},  6},
		{[]int{-5, 5}, 0},
		{[]int{}, 0},
	}
	for _, input := range dict {
		var actual = Sum(input.values)
		if actual != input.result {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%v"`, actual, input.result, input.values)
		}
	}
}
