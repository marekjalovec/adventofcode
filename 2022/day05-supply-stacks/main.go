package main

import (
	"fmt"
	"strings"

	"adventofcode/utils"
)

func getStacks() [][]string {
	return [][]string{
		{"B", "G", "S", "C"},
		{"T", "M", "W", "H", "J", "N", "V", "G"},
		{"M", "Q", "S"},
		{"B", "S", "L", "T", "W", "N", "M"},
		{"J", "Z", "F", "T", "V", "G", "W", "P"},
		{"C", "T", "B", "G", "Q", "H", "S"},
		{"T", "J", "P", "B", "W"},
		{"G", "D", "C", "Z", "F", "T", "Q", "M"},
		{"N", "S", "H", "B", "P", "F"},
	}
}

func moveCrate(n int, from int, to int, stacks [][]string) {
	batch := stacks[from][len(stacks[from])-n:]
	stacks[to] = append(stacks[to], batch...)
	stacks[from] = stacks[from][:len(stacks[from])-n]
}

func firstStar(lines []string, stacks [][]string) string {
	for _, line := range lines {
		l := strings.Split(line, " ")
		n, from, to := utils.ParseInt(l[0]), utils.ParseInt(l[1])-1, utils.ParseInt(l[2])-1
		for i := 0; i < n; i++ {
			moveCrate(1, from, to, stacks)
		}
	}
	return strings.Join(utils.Map(stacks, func(s []string) string {
		return s[len(s)-1]
	}), "")
}

func secondStar(lines []string, stacks [][]string) string {
	for _, line := range lines {
		l := strings.Split(line, " ")
		n, from, to := utils.ParseInt(l[0]), utils.ParseInt(l[1])-1, utils.ParseInt(l[2])-1
		moveCrate(n, from, to, stacks)
	}
	return strings.Join(utils.Map(stacks, func(s []string) string {
		return s[len(s)-1]
	}), "")
}

func main() {
	var lines = utils.ReadFile("2022/day05-supply-stacks/input.txt")

	fmt.Println("First Star:", firstStar(lines, getStacks()))
	fmt.Println("Second Star:", secondStar(lines, getStacks()))
}
