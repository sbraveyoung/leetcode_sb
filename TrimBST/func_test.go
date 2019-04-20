package trimBST

import (
	"reflect"
	"testing"
)

func Test_trimBST(t *testing.T) {
	type args struct {
		root *TreeNode
		L    int
		R    int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				L: 1,
				R: 2,
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			name: "second",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 0,
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 1,
							},
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				L: 1,
				R: 3,
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBST(tt.args.root, tt.args.L, tt.args.R); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
