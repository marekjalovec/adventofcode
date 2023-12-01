package main

import (
	"fmt"
	"sort"
	"strconv"

	"adventofcode/utils"
)

func getSumForTopN(lines []string, n int) int {
	var buff []int
	var sum int

	for _, v := range lines {
		if v == "" {
			buff = append(buff, sum)
			sum = 0
		} else {
			pv, _ := strconv.ParseInt(v, 10, 0)
			sum += int(pv)
		}
	}
	buff = append(buff, sum)

	sort.Sort(sort.Reverse(sort.IntSlice(buff)))

	return utils.Sum(buff[0:n])
}

func firstStar(lines []string) int {
	return getSumForTopN(lines, 1)
}

func secondStar(lines []string) int {
	return getSumForTopN(lines, 3)
}

func main() {
	var lines = utils.ReadFile("2022/day01-calorie-counting/input.txt")

	fmt.Println("First Star:", firstStar(lines))
	fmt.Println("Second Star:", secondStar(lines))
}
