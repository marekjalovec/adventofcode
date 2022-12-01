package utils

func Sum[T int](data []T) T {
	var res T

	for _, v := range data {
		res += v
	}

	return res
}
