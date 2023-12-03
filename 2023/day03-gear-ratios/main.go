package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"adventofcode/utils"
)

func getRune(lines *[]string, x int, y int) rune {
	if x < 0 || x > len((*lines)[0])-1 {
		return '.'
	}

	if y < 0 || y > len(*lines)-1 {
		return '.'
	}

	return rune((*lines)[y][x])
}

func getRange(lines *[]string, x int, x2 int, y int) []byte {
	x = utils.MinMax(x, 0, len((*lines)[0])-1)
	x2 = utils.MinMax(x2, 0, len((*lines)[0])-1)

	if y < 0 || y > len(*lines)-1 {
		return []byte("")
	}

	return []byte((*lines)[y][x:x2])
}

func isSymbol(lines *[]string, x int, y int) bool {
	r := getRune(lines, x, y)
	return !unicode.IsDigit(r) && r != '.'
}

func isStar(lines *[]string, x int, y int) bool {
	return getRune(lines, x, y) == '*'
}

func starCoords(x int, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func addSpoke(hm *map[string][]int, x int, y int, val int) {
	_, ok := (*hm)[starCoords(x, y)]
	if !ok {
		(*hm)[starCoords(x, y)] = []int{}
	}
	(*hm)[starCoords(x, y)] = append((*hm)[starCoords(x, y)], val)
}

func firstStar(lines []string) int {
	sum := 0
	r := regexp.MustCompile(`[^0-9.]`)
	for lN, line := range lines {
		var first, last = -1, -1
		for chrN, char := range line {
			if unicode.IsDigit(char) && first == -1 {
				first = chrN
			}
			if first != -1 {
				if chrN == len(line)-1 {
					last = chrN
				}
				if !unicode.IsDigit(char) {
					last = chrN - 1
				}
			}

			if first != -1 && last != -1 {
				if isSymbol(&lines, first-1, lN) ||
					isSymbol(&lines, last+1, lN) ||
					r.Match(getRange(&lines, first-1, last+2, lN-1)) ||
					r.Match(getRange(&lines, first-1, last+2, lN+1)) {
					sum += utils.ParseInt(line[first : last+1])
				}
				first = -1
				last = -1
			}
		}
	}

	return sum
}

func secondStar(lines []string) int {
	sum := 0
	hubMap := make(map[string][]int)
	for lN, _ := range lines {
		lines[lN] = fmt.Sprintf(".%s.", lines[lN])
	}
	for lN, line := range lines {
		var first, last = -1, -1
		for chrN, char := range line {
			if unicode.IsDigit(char) && first == -1 {
				first = chrN
			}
			if first != -1 {
				if chrN == len(line)-1 {
					last = chrN
				}
				if !unicode.IsDigit(char) {
					last = chrN - 1
				}
			}

			if first != -1 && last != -1 {
				num := utils.ParseInt(line[first : last+1])
				isStarBefore := isStar(&lines, first-1, lN)
				if isStarBefore {
					addSpoke(&hubMap, first-1, lN, num)
				}
				isStarAfter := isStar(&lines, last+1, lN)
				if isStarAfter {
					addSpoke(&hubMap, last+1, lN, num)
				}
				beforeRange := string(getRange(&lines, first-1, last+2, lN-1))
				if strings.Contains(beforeRange, "*") {
					i := strings.Index(beforeRange, "*")
					addSpoke(&hubMap, first+i-1, lN-1, num)
				}
				afterRange := string(getRange(&lines, first-1, last+2, lN+1))
				if strings.Contains(afterRange, "*") {
					i := strings.Index(afterRange, "*")
					addSpoke(&hubMap, first+i-1, lN+1, num)
				}

				first = -1
				last = -1
			}
		}
	}

	for _, ints := range hubMap {
		if len(ints) > 1 {
			sum += utils.Product(ints)
		}
	}

	return sum
}

func main() {
	var lines = utils.ReadFile("2023/day03-gear-ratios/input.txt")

	fmt.Println("First Star:", firstStar(lines))
	fmt.Println("Second Star:", secondStar(lines))
}
