package epsslice

import (
	"reflect"
	"testing"
)

func TestInt_IndexOf(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		g    Int
		args args
		want int
	}{
		{Int{1, -5, 7}, args{1}, 0},
		{Int{1, -5, 7}, args{-5}, 1},
		{Int{1, -5, 7, 7, 7}, args{7}, 2},
		{Int{1, -5, 7}, args{6}, -1},
		{Int{}, args{0}, -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.g.IndexOf(tt.args.i); got != tt.want {
				t.Errorf("Int.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Copy(t *testing.T) {
	tests := []struct {
		g    Int
		want Int
	}{
		{Int{}, Int{}},
		{Int{-5, 5}, Int{-5, 5}},
		{Int{1}, Int{1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.g.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Sum(t *testing.T) {
	tests := []struct {
		g    Int
		want int
	}{
		{Int{1, 5, 7}, 13},
		{Int{}, 0},
		{Int{1, -5, 7}, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.g.Sum(); got != tt.want {
				t.Errorf("Int.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_RemoveIndex(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		g    *Int
		g2   Int
		args args
		want bool
	}{
		{&Int{1, -5, 7}, Int{-5, 7}, args{0}, true},
		{&Int{1, -5, 7}, Int{1, 7}, args{1}, true},
		{&Int{1, -5, 7}, Int{1, -5}, args{2}, true},
		{&Int{1, -5, 7}, Int{1, -5, 7}, args{-1}, false},
		{&Int{1, -5, 7}, Int{1, -5, 7}, args{3}, false},
		{&Int{}, Int{}, args{3}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.g.RemoveIndex(tt.args.i)
			if got != tt.want || tt.g.Compare(tt.g2) != 0 {
				t.Errorf("Int.RemoveIndex() = %v (%v), want %v (%v)", got, tt.g, tt.want, tt.g2)
			}
		})
	}
}

func TestInt_Add(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		g    *Int
		args args
		g2   Int
	}{
		{&Int{1, -5, 7}, args{11}, Int{1, -5, 7, 11}},
		{&Int{1, -5, 7}, args{-11}, Int{1, -5, 7, -11}},
		{&Int{}, args{11}, Int{11}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tt.g.Add(tt.args.n)

			if tt.g.Compare(tt.g2) != 0 {
				t.Errorf("Int.Add() = (%v), want (%v)", tt.g, tt.g2)
			}
		})
	}
}

func TestInt_ElementClampPanic(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		g    Int
		args args
		want int
	}{
		{Int{5, 7, 10}, args{-1}, 5},
		{Int{5, 7, 10}, args{0}, 5},
		{Int{5, 7, 10}, args{1}, 7},
		{Int{5, 7, 10}, args{2}, 10},
		{Int{5, 7, 10}, args{3}, 10},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.g.ElementClampPanic(tt.args.i); got != tt.want {
				t.Errorf("Int.ElementClampPanic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Compare(t *testing.T) {
	type args struct {
		g2 Int
	}
	tests := []struct {
		g    Int
		args args
		want int
	}{
		{Int{}, args{Int{}}, 0},
		{Int{1, 5, 6}, args{Int{1, 5, 6}}, 0},
		{Int{1, 5, 6}, args{Int{1, 5, 9}}, -1},
		{Int{1, 5, 6}, args{Int{1, 5, -6}}, 1},

		{Int{}, args{Int{6}}, -1},
		{Int{}, args{Int{-6}}, 1},

		{Int{6}, args{Int{}}, 1},
		{Int{-6}, args{Int{}}, -1},

		{Int{5, -1, 7}, args{Int{5, -1, 7, 6}}, -1},
		{Int{5, -1, 7}, args{Int{5, -1, 7, -6}}, 1},
		{Int{5, -1, 7, 6}, args{Int{5, -1, 7}}, 1},
		{Int{5, -1, 7, -6}, args{Int{5, -1, 7}}, -1},

		{Int{}, args{Int{0}}, 0},
		{Int{0}, args{Int{}}, 0},
		{Int{5, -1, 7}, args{Int{5, -1, 7, 0}}, 0},
		{Int{5, -1, 7, 0}, args{Int{5, -1, 7}}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.g.Compare(tt.args.g2); got != tt.want {
				t.Errorf("Int.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Pop(t *testing.T) {
	tests := []struct {
		g       *Int
		g2      Int
		want    int
		wantErr bool
	}{
		{&Int{5, 6, 7}, Int{5, 6}, 7, false},
		{&Int{5, 6}, Int{5}, 6, false},
		{&Int{5}, Int{}, 5, false},
		{&Int{}, Int{}, 0, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := tt.g.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want || tt.g.Compare(tt.g2) != 0 {
				t.Errorf("Int.Pop() = %v (%v), want %v (%v)", got, tt.g, tt.want, tt.g2)
			}
		})
	}
}

func TestInt_Contains(t *testing.T) {
	tests := []struct {
		g    Int
		i    int
		want bool
	}{
		{g: Int{1, -5, 7}, i: 1, want: true},
		{g: Int{1, -5, 7}, i: -5, want: true},
		{g: Int{1, -5, 7}, i: 7, want: true},
		{g: Int{1, 5, 7}, i: 3, want: false},
		{g: Int{}, i: 3, want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.g.Contains(tt.i); got != tt.want {
				t.Errorf("Int.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Min(t *testing.T) {
	tests := []struct {
		g    Int
		want int
	}{
		{Int{1, 5, 7}, 1},
		{Int{}, 0},
		{Int{1, -5, 7}, -5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.g.Min(); got != tt.want {
				t.Errorf("Int.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
