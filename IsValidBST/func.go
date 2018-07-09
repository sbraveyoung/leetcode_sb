package isValidBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
// wrong
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return validBST(root.Left, root.Val, -2147483648, true) &&
		validBST(root.Right, 2147483647, root.Val, false)
}

func validBST(root *TreeNode, max, min int, isLeft bool) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	if isLeft {
		return validBST(root.Left, root.Val, -2147483648, true) && validBST(root.Right, max, root.Val, false)
	} else {
		return validBST(root.Left, root.Val, min, true) && validBST(root.Right, 2147483647, root.Val, false)
	}
}
*/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	result := []int{}
	stack := []*TreeNode{}
	p := root
	for len(stack) != 0 || p != nil {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, top.Val)
			p = top.Right
		}
	}
	return isSorted(result)
}

func isSorted(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return true
	}
	for i := 1; i < length; i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}
