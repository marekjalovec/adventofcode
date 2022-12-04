package utils

import (
	"strconv"
)

func ParseInteger(value string) int {
	var v, _ = strconv.ParseInt(value, 10, 0)

	return int(v)
}
