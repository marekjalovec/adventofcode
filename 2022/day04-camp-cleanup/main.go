package main

import (
	"fmt"
	"strings"

	"adventofcode/utils"
)

func firstStar(lines []string) int {
	var sum int
	for _, line := range lines {
		ranges := utils.Map(strings.Split(line, ","), func(s string) []int {
			return utils.Map(strings.Split(s, "-"), utils.ParseInt)
		})

		if (ranges[0][0] >= ranges[1][0] && ranges[0][1] <= ranges[1][1]) ||
			(ranges[1][0] >= ranges[0][0] && ranges[1][1] <= ranges[0][1]) {
			sum++
		}
	}

	return sum
}

func secondStar(lines []string) int {
	var sum int
	for _, line := range lines {
		ranges := utils.Map(strings.Split(line, ","), func(s string) []int {
			return utils.Map(strings.Split(s, "-"), utils.ParseInt)
		})

		if (ranges[0][0] >= ranges[1][0] && ranges[0][0] <= ranges[1][1]) ||
			(ranges[1][0] >= ranges[0][0] && ranges[1][0] <= ranges[0][1]) {
			sum++
		}
	}

	return sum
}

func main() {
	var lines = utils.ReadFile("2022/day04-camp-cleanup/input.txt")

	fmt.Println("First Star:", firstStar(lines))
	fmt.Println("Second Star:", secondStar(lines))
}
