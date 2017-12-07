package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	ret := &ListNode{}
	tmp := ret
	flow := false
	for l1 != nil && l2 != nil {
		ret.Next = &ListNode{}
		ret = ret.Next
		ret.Val = l1.Val + l2.Val
		if flow {
			ret.Val += 1
		}
		if ret.Val > 9 {
			flow = true
			ret.Val -= 10
		} else {
			flow = false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	var list *ListNode
	if l1 == nil {
		list = l2
	} else {
		list = l1
	}
	if list == nil {
		if flow {
			ret.Next = &ListNode{Val: 1}
		}
	}
	for list != nil {
		ret.Next = &ListNode{}
		ret = ret.Next
		if flow {
			ret.Val = list.Val + 1
		} else {
			ret.Val = list.Val
		}
		if ret.Val > 9 {
			flow = true
			ret.Val -= 10
		} else {
			flow = false
		}
		list = list.Next
	}
	if flow {
		ret.Next = &ListNode{Val: 1}
	}
	return tmp.Next
}
