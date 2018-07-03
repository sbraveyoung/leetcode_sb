package minStack

import (
	"testing"
)

func Test(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	if obj.GetMin() != -3 {
		t.Error("obj.GetMin()!=-3")
	}
	obj.Pop()
	if obj.Top() != 0 {
		t.Error("obj.Top()!=0")
	}
	if obj.GetMin() != -2 {
		t.Error("obj.GetMin()!=0")
	}
}
