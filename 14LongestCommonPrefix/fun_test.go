package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{
		"hello world",
		"good good study,day day up",
	}
	if longestCommonPrefix(strs) != "" {
		t.Errorf("Test longestCommonPrefix fail.")
	}
	strs = []string{
		"abstract",
		"abandon",
	}
	if longestCommonPrefix(strs) != "ab" {
		t.Errorf("Test longestCommonPrefix fail.")
	}
	strs = []string{
		"abcdefg",
		"abcdefg",
	}
	if longestCommonPrefix(strs) != "abcdefg" {
		t.Errorf("Test longestCommonPrefix fail.")
	}
	strs = []string{
		"abcdef",
		"abcdefg",
	}
	if longestCommonPrefix(strs) != "abcdef" {
		t.Errorf("Test longestCommonPrefix fail.")
	}
	strs = []string{
		"",
		"",
	}
	if longestCommonPrefix(strs) != "" {
		t.Errorf("Test longestCommonPrefix fail.")
	}
	strs = []string{
		"a",
		"a",
	}
	if longestCommonPrefix(strs) != "a" {
		t.Errorf("Test longestCommonPrefix fail.")
	}
}
