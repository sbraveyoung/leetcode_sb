package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

/*
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newList := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newList
}
*/

/*
func reverseList(head *ListNode) *ListNode {
	stack := []*ListNode{}
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}
	var tail *ListNode
label:
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		node.Next = nil
		stack = stack[:len(stack)-1]
		if head == nil {
			head = node
			tail = head
			continue label
		}
		tail.Next = node
		tail = tail.Next
	}
	return head
}
*/
