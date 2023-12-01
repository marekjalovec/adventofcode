package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"

	"adventofcode/utils"
)

func firstStar(lines []string) int {
	var sum int
	reg := regexp.MustCompile(`^[^0-9]*([0-9]).*?([0-9]?)[^0-9]*$`)
	for _, v := range lines {
		line := reg.FindStringSubmatch(v)
		var num int
		if line[2] != "" {
			num, _ = strconv.Atoi(fmt.Sprintf("%s%s", line[1], line[2]))
		} else {
			num, _ = strconv.Atoi(fmt.Sprintf("%s%s", line[1], line[1]))
		}
		sum += num
	}
	return sum
}

var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func secondStar(lines []string) int {
	sum := 0
	findNum := regexp.MustCompile(strings.Join(maps.Keys(numMap), "|"))
	for _, v := range lines {
		first := findNum.FindString(v)
		var last string
		for i := len(v); i > 0; i-- {
			last = findNum.FindString(v[i-1:])
			if last != "" {
				break
			}
		}
		num, _ := strconv.Atoi(fmt.Sprintf("%s%s", numMap[first], numMap[last]))
		sum += num
	}

	return sum
}

func main() {
	var lines = utils.ReadFile("2023/day01-calibration-document/input.txt")

	fmt.Println("First Star:", firstStar(lines))
	fmt.Println("Second Star:", secondStar(lines))
}
