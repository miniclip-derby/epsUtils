package epsSlice

import "github.com/miniclip-derby/epsUtils/epsInt64"

// Int64 is a slice of ints
type Int64 []int64

// IndexOf returns the index of the first occurance of s or -1 if not found
func (g Int64) IndexOf(i int64) int {
	for n, v := range g {
		if v == i {
			return n
		}
	}
	return -1
}

// ToStringSlice returns a string slice representation of this int64 slice
func (g Int64) ToStringSlice() String {
	s := String{}
	for _, v := range g {
		s = append(s, epsInt64.ToString(v))
	}
	return s
}

// RemoveIndex removes an element at index i
func (g *Int64) RemoveIndex(i int) bool {
	*g = append((*g)[:i], (*g)[i+1:]...)
	return true
}

// Add appends to the slice
func (g *Int64) Add(n int64) {
	*g = append(*g, n)
}

// ElementClampPanic returns the element at index i clamped to the slice.  If the slice is empty it panics
func (g Int64) ElementClampPanic(i int) int64 {
	if len(g) == 0 {
		panic("Slice is empty")
	}
	if i < 0 {
		return g[0]
	}
	if i >= len(g) {
		return g[len(g)-1]
	}

	return g[i]
}
