package findTarget

//https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/submissions/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	return find(root, root, k)
}

func find(root, cur *TreeNode, k int) bool {
	if root == nil || cur == nil {
		return false
	}
	sub := k - cur.Val
	if findSub(root, cur, sub) == true {
		return true
	}
	return find(root, cur.Left, k) || find(root, cur.Right, k)
}

func findSub(root, cur *TreeNode, sub int) bool {
	if root == nil || cur == nil {
		return false
	}
	if cur != root && root.Val == sub {
		return true
	} else if root.Val > sub {
		return findSub(root.Left, cur, sub)
	} else {
		return findSub(root.Right, cur, sub)
	}
}
