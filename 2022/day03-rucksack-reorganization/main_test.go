package main

import (
	"adventofcode/utils"
	"testing"
)

var lines = utils.ReadFile("input_test.txt")

type TestCase struct {
	inputs []string
	result int32
}

func TestFindDuplicates(t *testing.T) {
	var dict = []TestCase{
		{splitLine("vJrwpWtwJgWrhcsFMMfFFhFp"), 'p'},
		{splitLine("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"), 'L'},
		{splitLine("PmmdzqPrVvPwwTWBwg"), 'P'},
		{splitLine("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"), 'v'},
		{splitLine("ttgJtRGJQctTZtZT"), 't'},
		{splitLine("CrZsJsPPZsGzwwsLwLmpwMDw"), 's'},

		{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}, 'r'},
		{[]string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}, 'Z'},
	}
	for _, input := range dict {
		var actual = findDuplicate(input.inputs)
		if actual != input.result {
			t.Fatalf(`Actual output %d does not match the expected output %d for input "%s"`, actual, input.result, input.inputs)
		}
	}
}

func TestFirstStar(t *testing.T) {
	var expected = 157
	var actual = firstStar(lines)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}

func TestSecondStar(t *testing.T) {
	var expected = 70
	var actual = secondStar(lines)

	if actual != expected {
		t.Fatalf(`Actual output %d does not match the expected output %d`, actual, expected)
	}
}
