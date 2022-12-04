package utils

import (
	"strconv"
)

func ParseInt(value string) int {
	var v, err = strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0
	}

	return int(v)
}
