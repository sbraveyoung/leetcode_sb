package isSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return _isSymmetric(root.Left, root.Right)
}

func _isSymmetric(left, right *TreeNode) bool {
	// left == right == nil
	if left == nil && right == nil {
		return true
	}
	// (left == nil && right !=nil) || (left != nil && right == nil)
	if left == nil || right == nil {
		return false
	}
	// left == right != nil
	if left.Val != right.Val {
		return false
	}
	return _isSymmetric(left.Left, right.Right) && _isSymmetric(left.Right, right.Left)
}

/*
// time out
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) != 0 {
		allLeaf := true
		length := len(queue)
		for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
			if queue[i] != nil {
				allLeaf = false
			}
			if queue[i] == nil && queue[j] == nil {
				continue
			}
			if queue[i] == nil || queue[j] == nil {
				return false
			}
			if queue[i].Val != queue[j].Val {
				return false
			}
		}
		if !allLeaf {
			for i := 0; i < length; i++ {
				if queue[i] == nil {
					queue = append(queue, nil, nil)
				} else {
					queue = append(queue, queue[i].Left, queue[i].Right)
				}
			}
			queue = queue[length:]
		} else {
			return true
		}
	}
	return true
}
*/
