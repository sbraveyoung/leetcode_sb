package levelOrder

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
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
				root: &TreeNode{
					Val: 1,
				},
			},
			want: [][]int{{1}},
		},
		{
			name: "second",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			want: [][]int{
				{1},
				{2, 4},
				{3, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
