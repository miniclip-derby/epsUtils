package epsslice

import (
	"testing"
)

func TestFloat32_Sum(t *testing.T) {
	tests := []struct {
		g    Float32
		want float32
	}{
		{Float32{}, 0},
		{Float32{1, 2, 3, 4, 5}, 15},
		{Float32{-10, 20}, 10},
		{Float32{50}, 50},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.g.Sum(); got != tt.want {
				t.Errorf("Float32.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_Max(t *testing.T) {
	tests := []struct {
		g    Float32
		want float32
	}{
		{Float32{4, 10}, 10},
		{Float32{10, 10}, 10},
		{Float32{-1, 50}, 50},
		{Float32{}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.g.Max(); got != tt.want {
				t.Errorf("Float32.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32_Add(t *testing.T) {
	type args struct {
		n float32
	}
	tests := []struct {
		name string
		g    *Float32
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Add(tt.args.n)
		})
	}
}
