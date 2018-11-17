package generateMatrix

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	tests := []struct {
		name    string
		args    int
		wantRet [][]int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: 3,
			wantRet: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name:    "second",
			args:    0,
			wantRet: [][]int{},
		},
		{
			name:    "third",
			args:    1,
			wantRet: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := generateMatrix(tt.args); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("generateMatrix() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
