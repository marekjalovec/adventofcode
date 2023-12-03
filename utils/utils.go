package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func MinMax(val int, low int, high int) int {
	if val < low {
		return low
	}

	if val > high {
		return high
	}

	return val
}

func Product[T int](data []T) T {
	var res T

	for k, v := range data {
		if k == 0 {
			res = v
		} else {
			res *= v
		}
	}

	return res
}

func Map[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))

	for _, v := range data {
		res = append(res, f(v))
	}

	return res
}

func Sum[T int](data []T) T {
	var res T

	for _, v := range data {
		res += v
	}

	return res
}

func ParseInt(value string) int {
	var v, err = strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0
	}

	return int(v)
}

func ReadFile(fileName string) []string {
	var lines []string

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf(`Error reading file "%s": %s`, fileName, err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}
