package epsslice

import (
	"testing"
	"time"
)

func TestTime_RemoveIndex(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		g    *Time
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.RemoveIndex(tt.args.i); got != tt.want {
				t.Errorf("Time.RemoveIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_Add(t *testing.T) {
	type args struct {
		n time.Time
	}
	tests := []struct {
		name string
		g    *Time
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
