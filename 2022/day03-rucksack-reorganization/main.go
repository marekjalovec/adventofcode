package main

import (
	"fmt"
	"strings"

	"adventofcode/utils"
)

func findDuplicate(inputs []string) int32 {
	var all bool
	for _, c := range inputs[0] {
		all = true
		for i := 1; i < len(inputs); i++ {
			if !strings.Contains(inputs[i], string(c)) {
				all = false
			}
		}

		if all {
			return c
		}
	}

	return 0
}

// getPriority
// Lowercase item types a through z have priorities 1 through 26. -> ASCII 97-122
// Uppercase item types A through Z have priorities 27 through 52. -> ASCII 65-90
func getPriority(char int32) int {
	if char >= 97 && char <= 122 {
		char -= 96
	} else {
		char -= 38
	}

	return int(char)
}

func splitLine(input string) []string {
	return []string{
		input[0 : len(input)/2],
		input[len(input)/2:],
	}
}

func firstStar(lines []string) int {
	var sum int

	for _, line := range lines {
		sum += getPriority(findDuplicate(splitLine(line)))
	}

	return sum
}

func secondStar(lines []string) int {
	var sum int
	var triplet []string

	for i := 0; i < len(lines)/3; i++ {
		triplet = lines[i*3 : (i+1)*3]

		sum += getPriority(findDuplicate(triplet))
	}

	return sum
}

func main() {
	var lines = utils.ReadFile("2022/day03-rucksack-reorganization/input.txt")

	fmt.Println("First Star:", firstStar(lines))
	fmt.Println("Second Star:", secondStar(lines))
}
