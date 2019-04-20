package trimBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	newRoot := crop(root, L, R)
	return newRoot
}

func crop(root *TreeNode, L, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		return crop(root.Right, L, R)
	}
	if root.Val > R {
		return crop(root.Left, L, R)
	}
	root.Left = crop(root.Left, L, R)
	root.Right = crop(root.Right, L, R)
	return root
}
