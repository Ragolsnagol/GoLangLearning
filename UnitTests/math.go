package main

import (
	"fmt"
	"math"
	"strconv"
)

func AddPositive(x int, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, fmt.Errorf("both numbers must be non-negative")
	}
	return x + y, nil
}

func SqrtNumber(x float64) string {
	if x < 0 {
		if x == -1 {
			return "i"
		}
		return strconv.FormatFloat(math.Sqrt(-x), 'f', -1, 64) + "i"
	}
	return strconv.FormatFloat(math.Sqrt(x), 'f', -1, 64)
}
