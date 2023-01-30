package utils

func Max[T int | float64](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T int | float64](a T, b T) T {
	if a < b {
		return a
	}
	return b
}
