package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	if n == count {
		return head.Next
	}
	p, step := head, count-n
	for ; step != 1; step, p = step-1, p.Next {
	}
	p.Next = p.Next.Next
	return head
}
