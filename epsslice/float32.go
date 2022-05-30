package epsslice

import "github.com/miniclip-derby/epsUtils/epsFloat32"

// Float32 is a slice of ints
type Float32 []float32

// Sum returns the sum of all elements in the slice
func (g Float32) Sum() float32 {
	var s float32
	for _, v := range g {
		s += v
	}
	return s
}

// Max returns the max of all elements in the slice
func (g Float32) Max() float32 {
	s := float32(0)
	for _, v := range g {
		s = epsFloat32.Max(s, v)
	}
	return s
}

// Add appends to the slice
func (g *Float32) Add(n float32) {
	*g = append(*g, n)
}
