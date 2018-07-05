package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	var tail *ListNode
	for l1 != nil && l2 != nil {
		var node *ListNode
		if l1.Val < l2.Val {
			node = l1
			l1 = l1.Next
		} else {
			node = l2
			l2 = l2.Next
		}
		if head == nil {
			head = node
			tail = head
		} else {
			tail.Next = node
			tail = tail.Next
		}
	}
	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}
	return head
}
