package epsSlice

import "github.com/miniclip-derby/epsUtils/epsFloat64"

// Float64 is a slice of ints
type Float64 []float64

// Sum returns the sum of all elements in the slice
func (g Float64) Sum() float64 {
	var s float64
	for _, v := range g {
		s += v
	}
	return s
}

// Max returns the max of all elements in the slice
func (g Float64) Max() float64 {
	s := float64(0)
	for _, v := range g {
		s = epsFloat64.Max(s, v)
	}
	return s
}

// ElementClampPanic returns the element at index i clamped to the slice.  If the slice is empty it panics
func (g Float64) ElementClampPanic(i int) float64 {
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
