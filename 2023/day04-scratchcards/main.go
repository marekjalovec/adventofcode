package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"adventofcode/utils"
)

func firstStar(lines []string) int {
	sum := 0
	for _, line := range lines {
		l := strings.Split(line, " | ")
		wn := utils.Filter(strings.Split(l[0], " "), func(s string) bool { return s != "" })
		cn := utils.Filter(strings.Split(l[1], " "), func(s string) bool { return s != "" })
		ol := utils.Overlap(wn, cn)
		if len(ol) > 0 {
			sum += int(math.Pow(2, float64(len(ol)-1)))
		}
	}

	return sum
}

func secondStar(lines []string) int {
	m := make([]int, len(lines))
	for i := range m {
		m[i] = 1
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		l := strings.Split(line, " | ")
		wn := utils.Filter(strings.Split(l[0], " "), func(s string) bool { return s != "" })
		cn := utils.Filter(strings.Split(l[1], " "), func(s string) bool { return s != "" })
		ol := utils.Overlap(wn, cn)
		for j := 1; j <= len(ol); j++ {
			m[i+j] += m[i]
		}
	}

	return utils.Sum(m)
}

func main() {
	lines := utils.ReadFile("2023/day04-scratchcards/input.txt")
	r1 := regexp.MustCompile(` +`)
	r2 := regexp.MustCompile(`^Card (\d+):`)
	for k, _ := range lines {
		lines[k] = r1.ReplaceAllString(lines[k], " ")
		lines[k] = r2.ReplaceAllString(lines[k], "")
	}

	fmt.Println("First Star:", firstStar(lines))
	fmt.Println("Second Star:", secondStar(lines))
}
