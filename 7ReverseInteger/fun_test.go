package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	args := []int{0, 1, -1, 10, -10, 2147483647, -2147483648}
	rets := []int{0, 1, -1, 1, -1, 0, 0}
	for index, arg := range args {
		if rets[index] != reverse(arg) {
			t.Errorf("TestTwoSum fail.arg:%d\tret:%d\texpect:%d\n", arg, reverse(arg), rets[index])
		}
	}
}
