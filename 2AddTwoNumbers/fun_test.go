package main

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	list1 := &ListNode{Val: 2}
	list1.Next = &ListNode{Val: 4}
	list1.Next.Next = &ListNode{Val: 3}

	list2 := &ListNode{Val: 5}
	list2.Next = &ListNode{Val: 6}
	list2.Next.Next = &ListNode{Val: 4}

	list := addTwoNumbers(list1, list2)
	for list != nil {
		fmt.Printf("%d -> ", list.Val)
		list = list.Next
	}
	fmt.Println("nil")
}

func TestAddTwoNumbers1(t *testing.T) {
	list1 := &ListNode{Val: 5}

	list2 := &ListNode{Val: 5}

	list := addTwoNumbers(list1, list2)
	for list != nil {
		fmt.Printf("%d -> ", list.Val)
		list = list.Next
	}
	fmt.Println("nil")
}

func TestAddTwoNumbers2(t *testing.T) {
	list1 := &ListNode{Val: 1}

	list2 := &ListNode{Val: 9}
	list2.Next = &ListNode{Val: 9}

	list := addTwoNumbers(list1, list2)
	for list != nil {
		fmt.Printf("%d -> ", list.Val)
		list = list.Next
	}
	fmt.Println("nil")
}
