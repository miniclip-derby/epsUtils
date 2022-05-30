package epsslice

// String is a slice of strings //////////////////////////////////////////////
type String []string

// IndexOf returns the index of the first occurrence of s or -1 if not found
func (g String) IndexOf(s string) int {
	for i, v := range g {
		if v == s {
			return i
		}
	}
	return -1
}

// Contains returns true if the string appears in the slice
func (g String) Contains(s string) bool {
	return g.IndexOf(s) != -1
}

// Copy returns a copy of the slice
func (g String) Copy() String {
	s := make(String, len(g))
	copy(s, g)
	return s
}

// RemoveIndex removes an element at index i
func (g *String) RemoveIndex(i int) bool {
	*g = append((*g)[:i], (*g)[i+1:]...)
	return true
}

// Add adds a string to the slice
func (g *String) Add(n string) {
	*g = append(*g, n)
}

// Add adds a string to the slice if it doesn't exist
func (g *String) AddUnique(n string) int {
	i := g.IndexOf(n)
	if i != -1 {
		return i
	}
	*g = append(*g, n)
	return len(*g) - 1
}

// Remove removes the first occurrence of s
func (g *String) Remove(s string) bool {
	i := g.IndexOf(s)
	if i == -1 {
		return false
	}
	g.RemoveIndex(i)
	return true
}

// StringSliceIndexOf does something
func StringSliceIndexOf(slice []string, s string) int {
	for i, v := range slice {
		if v == s {
			return i
		}
	}
	return -1
}

// ElementClampPanic returns the element at index i clamped to the slice.  If the slice is empty it panics
func (g String) ElementClampPanic(i int) string {
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
