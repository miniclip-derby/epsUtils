package epsInt

import (
	"fmt"
	"math/rand"
	"strconv"
)

// ToString converts int to string
func ToString(i int) string {
	return strconv.Itoa(i)
}

// Rand returns rand [0-max)
func Rand(max int) int {
	return rand.Intn(max)
}

// RandRange returns a random number between [min and max)
func RandRange(min int, max int) int {
	return Rand(max-min) + min
}

// RandRangeSafe returns a random number between [min and max)
func RandRangeSafe(min int, max int) int {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return Rand(max-min) + min
}

// ToMoney returns a "£pounds.pence" string in of i pence
func ToMoney(i int) string {
	pounds := i / 100
	pence := i - (pounds * 100)
	return fmt.Sprintf("£%v.%02d", pounds, pence)
}

//	func int32Bucket(value int, bucket int) int {
//		return value - (value % bucket)
//	}
//
//	func intBucket(value int64, bucket int64) int64 {
//		return value - (value % bucket)
//	}
//
// Abs returns the abs of v
func Abs_deprecated(v int) int {
	if v < 0 {
		return -1
	} else if v > 0 {
		return 1
	}
	return 0
}

// Sign returns the sign of v
func Sign(v int) int {
	if v < 0 {
		return -1
	} else if v > 0 {
		return 1
	}
	return 0
}

// Abs2 returns the abs of v
func Abs2(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// Min returns min of v1 and v2
func Min(v1 int, v2 int) int {

	if v1 < v2 {
		return v1
	}
	return v2
}

// Max returns max of v1 and v2
func Max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

//
// func int64Max(v1 int64, v2 int64) int64 {
// 	if v1 > v2 {
// 		return v1
// 	}
// 	return v2
// }
//
//

// Limit limits v between min and max
func Limit(v int, min int, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

//
// func int64Limit(v int64, min int64, max int64) int64 {
// 	if v < min {
// 		return min
// 	}
// 	if v > max {
// 		return max
// 	}
// 	return v
// }
