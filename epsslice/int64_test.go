package epsslice

import (
	"reflect"
	"testing"
)

func TestInt64_IndexOf(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		g    Int64
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.IndexOf(tt.args.i); got != tt.want {
				t.Errorf("Int64.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_ToStringSlice(t *testing.T) {
	tests := []struct {
		name string
		g    Int64
		want String
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.ToStringSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64.ToStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_RemoveIndex(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		g    *Int64
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.RemoveIndex(tt.args.i); got != tt.want {
				t.Errorf("Int64.RemoveIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64_Add(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		g    *Int64
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

func TestInt64_ElementClampPanic(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		g    Int64
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.ElementClampPanic(tt.args.i); got != tt.want {
				t.Errorf("Int64.ElementClampPanic() = %v, want %v", got, tt.want)
			}
		})
	}
}
