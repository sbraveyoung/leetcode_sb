package threeEqualParts

import (
	"reflect"
	"testing"
)

func Test_threeEqualParts(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				A: []int{1, 0, 1, 0, 1},
			},
			want: []int{0, 3},
		},
		{
			name: "second",
			args: args{
				A: []int{1, 1, 0, 1, 1},
			},
			want: []int{-1, -1},
		},
		{
			name: "third",
			args: args{
				A: []int{0, 0},
			},
			want: []int{-1, -1},
		},
		{
			name: "fourth",
			args: args{
				A: []int{0, 0, 0},
			},
			want: []int{0, 2},
		},
		{
			name: "fifth",
			args: args{
				A: []int{1, 1, 1},
			},
			want: []int{0, 2},
		},
		{
			name: "sixth",
			args: args{
				A: []int{1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1},
			},
			want: []int{2, 8},
		},
		{
			name: "seventh",
			args: args{
				A: []int{0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1},
			},
			want: []int{5, 11},
		},
		{
			name: "eighth",
			args: args{
				A: []int{0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1},
			},
			want: []int{5, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeEqualParts(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeEqualParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
