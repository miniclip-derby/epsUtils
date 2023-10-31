package epsSlice

import (
	"fmt"

	"github.com/miniclip-derby/epsUtils/epsInt"
)

// Int is a slice of ints
type Int []int

// IndexOf returns the index of the first occurrence of s or -1 if not found
func (g Int) IndexOf(i int) int {
	for n, v := range g {
		if v == i {
			return n
		}
	}
	return -1
}

func (g Int) Contains(i int) bool {
	return g.IndexOf(i) != -1
}

// Copy returns a copy of the slice
func (g Int) Copy() Int {
	s := make(Int, len(g))
	copy(s, g)
	return s
}

// Min returns the smallest value in the slice
// If the slice is empty returns 0
func (g Int) Min() int {
	if len(g) == 0 {
		return 0
	}
	min := g[0]
	for _, v := range g {
		if v < min {
			min = v
		}
	}

	return min
}

// Sum returns the sum of all elements in the slice
func (g Int) Sum() int {
	var s int
	for _, v := range g {
		s += v
	}
	return s
}

// RemoveIndex removes an element at index i
func (g *Int) RemoveIndex(i int) bool {
	if i < 0 || i >= len(*g) {
		return false
	}
	*g = append((*g)[:i], (*g)[i+1:]...)
	return true
}

// Add appends to the slice
func (g *Int) Add(n int) {
	*g = append(*g, n)
}

// Pop removes and returns the last element
func (g *Int) Pop() (int, error) {
	if len(*g) == 0 {
		return 0, fmt.Errorf("Can't pop empty slice")
	}
	var x int
	x, *g = (*g)[len(*g)-1], (*g)[:len(*g)-1]
	return x, nil
}

// Compare compares the elements in g to g2 and returns 0 if they equal.
// If the slices are of different length then the missing index is assumed to be 0
// n = the index of the differing element
// if g[n] < g2[n] then returns -1
// if g[n] > g2[n] then returns 1
func (g Int) Compare(g2 Int) int {

	if len(g) < len(g2) {
		return -epsInt.Sign(g2[len(g)])
	}
	if len(g2) < len(g) {
		return epsInt.Sign(g[len(g2)])
	}

	for i := range g {
		v := g[i] - g2[i]
		if v != 0 {
			return epsInt.Sign(v)
		}
	}
	return 0
}

// ElementClampPanic returns the element at index i clamped to the slice.  If the slice is empty it panics
func (g Int) ElementClampPanic(i int) int {
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
