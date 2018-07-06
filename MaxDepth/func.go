package maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	max := left
	if max < right {
		max = right
	}
	return max + 1
}

/*
func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) != 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			front := queue[i]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
		queue = queue[size:]
	}
	return depth
}
*/
