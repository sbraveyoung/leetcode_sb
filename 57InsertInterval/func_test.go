package insertPackage

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				intervals: [][]int{
					{1, 3},
					{6, 9},
				},
				newInterval: []int{2, 5},
			},
			want: [][]int{
				{1, 5},
				{6, 9},
			},
		},
		{
			name: "second",
			args: args{
				intervals: [][]int{
					{1, 2},
					{3, 5},
					{6, 7},
					{8, 10},
					{12, 16},
				},
				newInterval: []int{4, 8},
			},
			want: [][]int{
				{1, 2},
				{3, 10},
				{12, 16},
			},
		},
		{
			name: "third",
			args: args{
				intervals: [][]int{
					{6, 9},
				},
				newInterval: []int{2, 5},
			},
			want: [][]int{
				{2, 5},
				{6, 9},
			},
		},
		{
			name: "fourth",
			args: args{
				intervals: [][]int{
					{1, 3},
				},
				newInterval: []int{5, 8},
			},
			want: [][]int{
				{1, 3},
				{5, 8},
			},
		},
		{
			name: "fifth",
			args: args{
				intervals: [][]int{
					{6, 9},
				},
				newInterval: []int{2, 7},
			},
			want: [][]int{
				{2, 9},
			},
		},
		{
			name: "sixth",
			args: args{
				intervals: [][]int{
					{1, 3},
				},
				newInterval: []int{2, 5},
			},
			want: [][]int{
				{1, 5},
			},
		},
		{
			name: "seventh",
			args: args{
				intervals: [][]int{
					{1, 3},
					{6, 9},
				},
				newInterval: []int{4, 5},
			},
			want: [][]int{
				{1, 3},
				{4, 5},
				{6, 9},
			},
		},
		{
			name: "eighth",
			args: args{
				intervals: [][]int{
					{1, 3},
					{6, 9},
				},
				newInterval: []int{0, 10},
			},
			want: [][]int{
				{0, 10},
			},
		},
		{
			name: "ninth",
			args: args{
				intervals:   [][]int{},
				newInterval: []int{0, 10},
			},
			want: [][]int{
				{0, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
