package main

import (
	"fmt"
	"strings"

	"adventofcode/utils"
)

func firstStar(lines []string) int {
	sum := 0
	for _, v := range lines {
		l := strings.Split(v, ": ")
		game := utils.ParseInt(strings.Replace(l[0], "Game ", "", -1))
		draws := strings.Split(l[1], "; ")
		valid := true
		for _, d := range draws {
			cubes := strings.Split(d, ", ")
			for _, c := range cubes {
				cx := strings.Split(c, " ")
				if cx[1] == "red" && utils.ParseInt(cx[0]) > 12 {
					valid = false
				} else if cx[1] == "green" && utils.ParseInt(cx[0]) > 13 {
					valid = false
				} else if cx[1] == "blue" && utils.ParseInt(cx[0]) > 14 {
					valid = false
				}
			}
		}
		if valid {
			sum += game
		}
	}

	return sum
}

func secondStar(lines []string) int {
	sum := 0
	for _, v := range lines {
		l := strings.Split(v, ": ")
		draws := strings.Split(l[1], "; ")
		red := 0
		green := 0
		blue := 0
		for _, d := range draws {
			cubes := strings.Split(d, ", ")
			for _, c := range cubes {
				cx := strings.Split(c, " ")
				if cx[1] == "red" && utils.ParseInt(cx[0]) > red {
					red = utils.ParseInt(cx[0])
				} else if cx[1] == "green" && utils.ParseInt(cx[0]) > green {
					green = utils.ParseInt(cx[0])
				} else if cx[1] == "blue" && utils.ParseInt(cx[0]) > blue {
					blue = utils.ParseInt(cx[0])
				}
			}
		}
		sum += red * green * blue
	}

	return sum
}

func main() {
	var lines = utils.ReadFile("2023/day02-cube-conundrum/input.txt")

	fmt.Println("First Star:", firstStar(lines))
	fmt.Println("Second Star:", secondStar(lines))
}
