package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	count := 0
	for node := head; node != nil; node = node.Next {
		count++
	}
	var next, prev *ListNode
	for i := 0; i < count/2; i++ {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	if count%2 != 0 {
		head = head.Next
	}
	for prev != nil && head != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}
	if prev == nil && head == nil {
		return true
	}
	return false
}
