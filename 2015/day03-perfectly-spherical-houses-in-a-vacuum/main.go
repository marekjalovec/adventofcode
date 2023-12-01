package main

import (
	"fmt"

	"adventofcode/utils"
)

type Coords struct {
	x int
	y int
}

func moveCoords(char int32, c Coords) Coords {
	switch char {
	case '^':
		c.y++
	case '>':
		c.x++
	case 'v':
		c.y--
	case '<':
		c.x--
	}

	return c
}

func firstStar(input string) int {
	var h = make(map[string]int)
	h["0x0"] = 1
	var ns = Coords{0, 0}

	for _, v := range input {
		ns = moveCoords(v, ns)
		h[fmt.Sprintf(`%dx%d`, ns.x, ns.y)]++
	}

	return len(h)
}

func secondStar(input string) int {
	var h = make(map[string]int)
	h["0x0"] = 2
	var ns = Coords{0, 0}
	var rs = Coords{0, 0}
	var c = 0
	var k Coords

	for _, v := range input {
		if c%2 == 0 {
			ns = moveCoords(v, ns)
			k = ns
		} else {
			rs = moveCoords(v, rs)
			k = rs
		}

		h[fmt.Sprintf(`%dx%d`, k.x, k.y)]++
		c++
	}

	return len(h)
}

func main() {
	var lines = utils.ReadFile("2015/day03-perfectly-spherical-houses-in-a-vacuum/input.txt")

	fmt.Println("First Star:", firstStar(lines[0]))
	fmt.Println("Second Star:", secondStar(lines[0]))
}
