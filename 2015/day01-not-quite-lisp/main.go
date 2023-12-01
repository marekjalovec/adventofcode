package main

import (
	"fmt"
	"strings"

	"adventofcode/utils"
)

func firstStar(input string) int {
	var up = strings.Count(input, "(")
	var down = strings.Count(input, ")")

	return up - down
}

func secondStar(input string) int {
	var counter = 0
	for index, char := range input {
		if char == '(' {
			counter++
		} else {
			counter--
		}

		if counter < 0 {
			return index + 1
		}
	}

	return -1
}

func main() {
	var lines = utils.ReadFile("2015/day01-not-quite-lisp/input.txt")

	fmt.Println("First Star:", firstStar(lines[0]))
	fmt.Println("Second Star:", secondStar(lines[0]))
}
