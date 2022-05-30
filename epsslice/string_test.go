package epsslice

import (
	"reflect"
	"testing"
)

func TestString_IndexOf(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		g    String
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.IndexOf(tt.args.s); got != tt.want {
				t.Errorf("String.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Copy(t *testing.T) {
	tests := []struct {
		name string
		g    String
		want String
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_RemoveIndex(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		g    *String
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.RemoveIndex(tt.args.i); got != tt.want {
				t.Errorf("String.RemoveIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_Add(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		g    *String
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

func TestString_Remove(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		g    *String
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Remove(tt.args.s); got != tt.want {
				t.Errorf("String.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSliceIndexOf(t *testing.T) {
	type args struct {
		slice []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSliceIndexOf(tt.args.slice, tt.args.s); got != tt.want {
				t.Errorf("StringSliceIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_AddUnique(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		g    *String
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.AddUnique(tt.args.n); got != tt.want {
				t.Errorf("String.AddUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
