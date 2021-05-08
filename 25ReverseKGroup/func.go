package reverseKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	if head == nil || k <= 1 {
		return head
	}

	var retHead, retTail *ListNode
out:
	for head != nil {
		start, end := head, head

		for count := 0; count < k; count++ {
			if end == nil {
				if retTail == nil {
					return head
				}
				retTail.Next = head
				break out
			}
			end = end.Next
		}
		head = end

		tmpListHead, tmpListTail := reverse(start, end)
		if retHead == nil {
			retHead, retTail = tmpListHead, tmpListTail
		} else {
			retTail.Next = tmpListHead
			retTail = tmpListTail
		}
	}
	return retHead
}

func reverse(start, end *ListNode) (head, tail *ListNode) {
	if start == nil {
		return nil, nil
	}
	if start.Next == end {
		start.Next = nil
		return start, start
	}

	retHead, retTail := reverse(start.Next, end)
	retTail.Next = start
	start.Next = nil
	retTail = start
	return retHead, retTail
}
