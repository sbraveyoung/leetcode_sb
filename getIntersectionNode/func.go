package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	var head *ListNode
	for head = headA; head != nil; head = head.Next {
		lenA++
	}
	for head = headB; head != nil; head = head.Next {
		lenB++
	}
	other := head
	sub := 0
	if lenA > lenB {
		head = headA
		other = headB
		sub = lenA - lenB
	} else {
		head = headB
		other = headA
		sub = lenB - lenA
	}
	for ; sub > 0; sub-- {
		head = head.Next
	}
	for ; head != nil && other != nil; head, other = head.Next, other.Next {
		if head == other {
			return head
		}
	}
	return nil
}
