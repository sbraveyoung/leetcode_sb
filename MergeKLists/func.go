package mergeKLists

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type arrays []*ListNode

func (a arrays) Len() int {
	return len(a)
}

func (a arrays) Less(i, j int) bool {
	if a[i] == nil {
		return false
	}
	if a[j] == nil {
		return true
	}
	if a[i].Val > a[j].Val {
		return false
	}
	return true
}

func (a arrays) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func mergeKLists(lists []*ListNode) *ListNode {
	arr := arrays(lists)
	var ret, cur *ListNode
	for {
		if len(arr) == 0 {
			break
		}
		sort.Sort(arr)
		if arr[0] == nil {
			arr = arr[1:]
			continue
		}
		if ret == nil {
			ret = arr[0]
			cur = ret
		} else {
			cur.Next = arr[0]
			cur = cur.Next
		}
		arr[0] = arr[0].Next
	}
	return ret
}
