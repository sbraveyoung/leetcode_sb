package intervalIntersection

import (
	"reflect"
	"testing"
)

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		A [][]int
		B [][]int
	}
	tests := []struct {
		name  string
		args  args
		wantC [][]int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				A: [][]int{
					{0, 2}, {5, 10}, {13, 23}, {24, 25},
				},
				B: [][]int{
					{1, 5}, {8, 12}, {15, 24}, {25, 26},
				},
			},
			wantC: [][]int{
				{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25},
			},
		},
		{
			name: "second",
			args: args{
				A: [][]int{
					{0, 2}, {4, 6}, {8, 10}, {12, 14},
				},
				B: [][]int{
					{2, 4}, {6, 8}, {10, 12}, {14, 16},
				},
			},
			wantC: [][]int{
				{2, 2}, {4, 4}, {6, 6}, {8, 8}, {10, 10}, {12, 12}, {14, 14},
			},
		},
		{
			name: "third",
			args: args{
				A: [][]int{
					{0, 2}, {5, 10}, {13, 23}, {24, 25},
				},
				B: [][]int{
					{3, 4}, {11, 12},
				},
			},
			wantC: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := intervalIntersection(tt.args.A, tt.args.B); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("intervalIntersection() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
