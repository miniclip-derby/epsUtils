package epsslice

import "testing"

func TestFloat64_Sum(t *testing.T) {
	tests := []struct {
		name string
		g    Float64
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Sum(); got != tt.want {
				t.Errorf("Float64.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64_Max(t *testing.T) {
	tests := []struct {
		name string
		g    Float64
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Max(); got != tt.want {
				t.Errorf("Float64.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
