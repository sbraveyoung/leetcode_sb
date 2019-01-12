package addTwoNumbers

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil {
		cur := l1
		s1 := ""
		for cur != nil {
			s1 = strconv.Itoa(cur.Val) + s1
			cur = cur.Next
		}
		fmt.Print(s1)
	}
	if l2 != nil {
		cur := l2
		s2 := ""
		for cur != nil {
			s2 = strconv.Itoa(cur.Val) + s2
			cur = cur.Next
		}
		fmt.Print(" + ", s2, " = ")
	}
	var ret *ListNode
	var cur *ListNode
	carry := false
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val
		if carry {
			sum += 1
			carry = false
		}
		if ret == nil {
			ret = &ListNode{Val: sum % 10}
			cur = ret
		} else {
			cur.Next = &ListNode{Val: sum % 10}
			cur = cur.Next
		}
		if sum > 9 {
			carry = true
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil || l2 != nil {
		if l2 != nil {
			l1 = l2
		}
		cur.Next = l1
		// cur = cur.Next
	}
	for carry || cur.Next != nil {
		if carry {
			if cur.Next == nil {
				cur.Next = &ListNode{}
			}
			cur.Next.Val += 1
			carry = false
		}
		if cur.Next.Val > 9 {
			carry = true
			cur.Next.Val %= 10
		}
		cur = cur.Next
	}
	if ret != nil {
		cur := ret
		s := ""
		for cur != nil {
			s = strconv.Itoa(cur.Val) + s
			cur = cur.Next
		}
		fmt.Println(s)
	}
	return ret
}
