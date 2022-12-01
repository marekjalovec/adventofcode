package main

import (
	"adventofcode/utils"
	"fmt"
	"sort"
	"strconv"
)

func getSumForTopN(fileName string, n int) int {
	var buff []int
	var sum int

	var lines = utils.ReadFile(fileName)
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

func firstStar() int {
	return getSumForTopN("input.txt", 1)
}

func secondStar() int {
	return getSumForTopN("input.txt", 3)
}

func main() {
	fmt.Println("First Star:", firstStar())
	fmt.Println("Second Star:", secondStar())
}
