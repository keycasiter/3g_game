package util

import (
	"fmt"
	"strconv"
)

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

func Float64Remain(num float64) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", num), 64)
	return value
}
