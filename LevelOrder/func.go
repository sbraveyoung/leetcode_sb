package levelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		level := []int{}
		for i := 0; i < length; i++ {
			if queue[i] == nil {
				continue
			}
			level = append(level, queue[i].Val)
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
		if len(level) != 0 {
			res = append(res, level)
		}
		queue = queue[length:]
	}
	return res
}
