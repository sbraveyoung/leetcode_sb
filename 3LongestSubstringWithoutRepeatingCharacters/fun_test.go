package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	strings := []string{"abcabcbb", "bbbb", "pwwkew", "abba"}
	res := []int{3, 1, 3, 2}
	for index, str := range strings {
		if lengthOfLongestSubstring(str) != res[index] {
			t.Errorf("Test fail.")
		}
	}
}
