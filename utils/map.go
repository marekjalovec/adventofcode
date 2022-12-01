package utils

func Map[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))

	for _, v := range data {
		res = append(res, f(v))
	}

	return res
}
