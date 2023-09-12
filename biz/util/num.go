package util

func DivFloat64(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	return x / y
}

func DivInt64(x, y int64) int64 {
	if y == 0 {
		return 0
	}
	return x / y
}
