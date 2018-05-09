package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	i, j := 0, len(str)-1
	for i < j {
		if str[i] == str[j] {
			i++
			j--
			continue
		}
		return false
	}
	return true
}
