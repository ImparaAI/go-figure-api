package util

import (
	"github.com/google/go-cmp/cmp"
	"math"
)

func FloatCompare(x float64, y float64, tolerance float64) int {
	floatCompare := getFloatEqualityComparer(tolerance)

	if cmp.Equal(x, y, floatCompare) {
		return 0
	}

	if x > y {
		return 1
	}

	return -1
}

func getFloatEqualityComparer(tolerance float64) cmp.Option {
	return cmp.Comparer(func(x, y float64) bool {
		diff := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0

		if math.IsNaN(diff / mean) {
			return true
		}

		return (diff / mean) < tolerance
	})
}
