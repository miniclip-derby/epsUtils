package epsString

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Force kick

// ToInt string to int
func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ToIntPanic panics if the string can't be parsed
func ToIntPanic(s string) int {
	value, err := ToInt(s)
	if err != nil {
		panic(err.Error())
	}
	return value
}

// ToInt0 string to int or 0 if error
func ToInt0(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return value
}

// ToInt64 string to int64
func ToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// ToInt640 string to int64
func ToInt640(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// ToFloat32 string to float32
func ToFloat32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	return float32(f), err
}

// ToFloat320 string to float32 or 0 if error
func ToFloat320(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}
	return float32(f)
}

// ToFloat64 string to float32
func ToFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	return f, err
}

// ToFloat640 string to float64 or 0 if error
func ToFloat640(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

// ToBool calls ParseBool
func ToBool(s string) (bool, error) {
	value, err := strconv.ParseBool(s)
	return value, err
}

// ToBool calls ParseBool
func ToBool0(s string) bool {
	value, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return value
}

// ToIntSlice separates the string by sep then returns a slice of ints
func ToIntSlice(s, sep string) ([]int, error) {
	ret := []int{}
	split := strings.Split(s, sep)
	for _, v := range split {
		val, err := ToInt(v)
		if err != nil {
			return nil, err
		}
		ret = append(ret, val)
	}
	return ret, nil
}

// ToFloat32Slice separates the string by sep then returns a slice of float32s
func ToFloat32Slice(s, sep string) ([]float32, error) {
	ret := []float32{}
	split := strings.Split(s, sep)
	for _, v := range split {
		val, err := ToFloat32(v)
		if err != nil {
			return nil, err
		}
		ret = append(ret, val)
	}
	return ret, nil
}

// ToFloat320Slice separates the string by sep then returns a slice of float32s
func ToFloat320Slice(s, sep string) []float32 {
	ret := []float32{}
	split := strings.Split(s, sep)
	for _, v := range split {
		ret = append(ret, ToFloat320(v))
	}
	return ret
}

// SecondsToTime0 converts a string of seconds to time.Time
func SecondsToTime0(secs string) time.Time {
	return time.Unix(ToInt640(secs), 0)
}

// FromMoney string of money to int pence
func FromMoney(s string) (int, error) {
	split := strings.Split(s, ".")
	if len(split) != 2 {
		return 0, fmt.Errorf("Bad price (.)")
	}
	if strings.HasPrefix(split[0], "£") == false {
		return 0, fmt.Errorf("Bad price (&)")
	}
	pounds, err := ToInt(strings.TrimPrefix(split[0], "£"))
	if err != nil {
		return 0, err
	}
	pence, err := ToInt(split[1])
	if err != nil {
		return 0, err
	}
	return (pounds * 100) + pence, nil
}
