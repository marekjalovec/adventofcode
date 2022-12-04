package main

import (
	"adventofcode/utils"
	"fmt"
	"sort"
	"strings"
)

func parseValues(input string) (int, int, int) {
	var values = utils.Map(strings.Split(input, "x"), utils.ParseInt)

	return values[0], values[1], values[2]
}

func paperForBox(input string) int {
	var a, b, c = parseValues(input)
	var ab = a * b
	var bc = b * c
	var ac = a * c
	var arr = []int{ab, bc, ac}
	sort.Ints(arr)

	return ab*2 + bc*2 + ac*2 + arr[0]
}

func ribbonForBox(input string) int {
	var a, b, c = parseValues(input)
	var ab = (a + b) * 2
	var bc = (b + c) * 2
	var ac = (a + c) * 2
	var arr = []int{ab, bc, ac}
	sort.Ints(arr)

	return arr[0] + (a * b * c)
}

func firstStar(input []string) int {
	var sum int

	for _, v := range input {
		sum += paperForBox(v)
	}

	return sum
}

func secondStar(input []string) int {
	var sum int

	for _, v := range input {
		sum += ribbonForBox(v)
	}

	return sum
}

func main() {
	var lines = utils.ReadFile("2015/day02-i-was-told-there-would-be-no-math/input.txt")

	fmt.Println("First Star:", firstStar(lines))
	fmt.Println("Second Star:", secondStar(lines))
}
