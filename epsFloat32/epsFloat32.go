package epsFloat32

import (
	"fmt"
	"math"
	"math/rand"
)

///////////////////////////////////////////////////////////////////////////////////
// Limit functions

// Max returns the max of v1 and v1
func Max(v1 float32, v2 float32) float32 {
	if v1 > v2 {
		return v1
	}
	return v2
}

// Min returns the in of v1 and v1
func Min(v1 float32, v2 float32) float32 {
	if v1 < v2 {
		return v1
	}
	return v2
}

// Limit limits v in the range min,max
func Limit(v float32, min float32, max float32) float32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// Abs returns the absolute value of v
func Abs(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

///////////////////////////////////////////////////////////////////////////////////
// Rounding functions

// CeilToInt is a function
func CeilToInt(x float32) int {
	return int(math.Ceil(float64(x)))
}

// CeilToInt64 is a function
func CeilToInt64(x float32) int64 {
	return int64(math.Ceil(float64(x)))
}

// FloorToInt is a function
func FloorToInt(x float32) int {
	return int(math.Floor(float64(x)))
}

// FloorToInt64 is a function
func FloorToInt64(x float32) int64 {
	return int64(math.Floor(float64(x)))
}

// RoundToInt is a function
func RoundToInt(x float32) int {
	return FloorToInt(x + 0.5)
}

// RoundToInt64 rounds x
func RoundToInt64(x float32) int64 {
	return FloorToInt64(x + 0.5)
}

///////////////////////////////////////////////////////////////////////////////////
// Random functions

// Rand return rand [0,max)
func Rand(max float32) float32 {
	return rand.Float32() * max
}

// RandRange return rand [min,max)
func RandRange(min float32, max float32) float32 {
	return Rand(max-min) + min
}

///////////////////////////////////////////////////////////////////////////////////
// String functions

// ToString is a function
func ToString(x float32) string {
	return fmt.Sprintf("%v", x)
}

///////////////////////////////////////////////////////////////////////////////////
// Fractions

// Fraction is a function
func Fraction(min, max, x float32) float32 {
	return (x - min) / (max - min)
}

// ExpandFraction is a function
func ExpandFraction(min, max, x float32) float32 {
	return ((max - min) * x) + min
}

// RangeValue is a function
func RangeValue(min, max, x, low, high float32) float32 {
	return ExpandFraction(low, high, Fraction(min, max, x))
}
