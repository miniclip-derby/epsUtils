package epsFloat64

import (
	"fmt"
	"math"
)

// Max returns the max of v1 and v1
func Max(v1 float64, v2 float64) float64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

// Limit limits v in the range min,max
func Limit(v float64, min float64, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// FloorToInt returns Floor(x)
func FloorToInt(x float64) int {
	return int(math.Floor(x))
}

// FloorToInt64 returns Floor(x)
func FloorToInt64(x float64) int64 {
	return int64(math.Floor(x))
}

// CeilToInt returns Ceil(x)
func CeilToInt(x float64) int {
	return int(math.Ceil(x))
}

// CeilToInt64 returns Ceil(x)
func CeilToInt64(value float64) int64 {
	return int64(math.Ceil(value))
}

// Round rounds x
func Round(x float64) float64 {
	return math.Floor(x + 0.5)
}

// RoundToInt rounds x
func RoundToInt(x float64) int {
	return FloorToInt(x + 0.5)
}

// RoundToInt64 rounds x
func RoundToInt64(x float64) int64 {
	return FloorToInt64(x + 0.5)
}

// Abs returns Abs(x)
func Abs(x float64) float64 {
	return math.Abs(x)
}

// ToString is a function
func ToString(x float64) string {
	return fmt.Sprintf("%v", x)
}
