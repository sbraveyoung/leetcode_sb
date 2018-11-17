package spiralOrder

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantRet: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "second",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
				},
			},
			wantRet: []int{1, 2, 3},
		},
		{
			name: "third",
			args: args{
				matrix: [][]int{
					{1},
					{4},
					{7},
				},
			},
			wantRet: []int{1, 4, 7},
		},
		{
			name: "fourth",
			args: args{
				matrix: [][]int{
					{1},
				},
			},
			wantRet: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := spiralOrder(tt.args.matrix); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("spiralOrder() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
