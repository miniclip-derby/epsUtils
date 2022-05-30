package epsInt64

import (
	"math/rand"
	"strconv"
)

func ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func ToPosString(n int64) string {
	str := ToString(n)

	if n >= 20 {
		n -= ((n / 20) * 20)
	}

	switch n {
	case 1:
		return str + "st"
	case 2:
		return str + "nd"
	case 3:
		return str + "rd"
	case 4, 5, 6, 7, 8, 9, 0, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
		return str + "th"
	}
	return "unknown " + str
}

func Min(v1 int64, v2 int64) int64 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func Max(v1 int64, v2 int64) int64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

// Limit limits v in the range min,max
func Limit(v int64, min int64, max int64) int64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func Abs(v int64) int64 {
	if v > 0 {
		return v
	}
	return -v
}

// Rand returns rand [0-max)
func Rand(max int64) int64 {
	return rand.Int63n(max)
}
