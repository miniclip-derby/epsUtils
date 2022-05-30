package epsslice

import "testing"

func TestBool_RemoveIndex(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		g    *Bool
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.RemoveIndex(tt.args.i); got != tt.want {
				t.Errorf("Bool.RemoveIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_Add(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		g    *Bool
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Add(tt.args.b)
		})
	}
}
