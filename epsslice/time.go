package epsslice

import "time"

// Time is a slice of time.Times
type Time []time.Time

// RemoveIndex removes an element at index i
func (g *Time) RemoveIndex(i int) bool {
	*g = append((*g)[:i], (*g)[i+1:]...)
	return true
}

// Add appends to the slice
func (g *Time) Add(n time.Time) {
	*g = append(*g, n)
}
