package sortedArrayToBST

import (
	"math"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
		},
		{
			name: "second",
			args: args{
				nums: []int{0, 5, 9},
			},
		},
		{
			name: "third",
			args: args{
				nums: []int{0, 5},
			},
		},
		{
			name: "fourth",
			args: args{
				nums: []int{0, 1, 2, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !isValidBST(got) {
				t.Errorf("sortedArrayToBST() = %v", got)
			}
		})
	}
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return false
	}
	if root.Left != nil && root.Right != nil {
		if root.Left.Val > root.Right.Val {
			return false
		}
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	max := left
	if right > max {
		max = right
	}
	return max + 1
}
