package main

import (
	"adventofcode/utils"
	"crypto/md5"
	"fmt"
)

func getSuffix(input string, prefixMatch string) int {
	var m []byte
	var h string
	for i := 0; i < 10000000; i++ {
		m = []byte(fmt.Sprintf(`%s%d`, input, i))
		h = fmt.Sprintf(`%x`, md5.Sum(m))

		if h[0:len(prefixMatch)] == prefixMatch {
			return i
		}
	}

	return -1
}

func firstStar(input string) int {
	return getSuffix(input, "00000")
}

func secondStar(input string) int {
	return getSuffix(input, "000000")

}

func main() {
	var lines = utils.ReadFile("2015/day04-the-ideal-stocking-stuffer/input.txt")

	fmt.Println("First Star:", firstStar(lines[0]))
	fmt.Println("Second Star:", secondStar(lines[0]))
}
