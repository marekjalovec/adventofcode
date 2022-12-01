package main

import (
	"adventofcode/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseValues(input string) (int, int, int) {
	var values = strings.Split(input, "x")
	var intVal = utils.Map(values, func(value string) int {
		var v, _ = strconv.ParseInt(value, 10, 0)

		return int(v)
	})

	return intVal[0], intVal[1], intVal[2]
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
	var lines = utils.ReadFile("2015/day02/input.txt")

	fmt.Println("First Star:", firstStar(lines))
	fmt.Println("Second Star:", secondStar(lines))
}
