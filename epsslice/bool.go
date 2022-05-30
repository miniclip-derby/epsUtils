package epsslice

// Bool is a slice of bools
type Bool []bool

// RemoveIndex removes an element at index i
func (g *Bool) RemoveIndex(i int) bool {
	*g = append((*g)[:i], (*g)[i+1:]...)
	return true
}

// Add appends to the slice
func (g *Bool) Add(b bool) {
	*g = append(*g, b)
}
