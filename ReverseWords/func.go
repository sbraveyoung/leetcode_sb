package reverseWords

import "errors"

func reverseWords(s string) string {
	bs := []byte(s)
	//reverseStr(bs, 0, len(s)-1)

	index := 0
	for {
		before := index
		index = findSpace(bs, index)
		reverseStr(bs, before, index-1)
		if index == -1 || index == len(bs) {
			break
		}
		index++
	}
	return string(bs)
}

func findSpace(bs []byte, index int) int {
	if index < 0 || index > len(bs) {
		return -1
	}
	for index < len(bs) {
		if bs[index] == ' ' {
			return index
		}
		index++
	}
	return len(bs)
}

func reverseStr(s []byte, left, right int) error {
	if left > right || left < 0 || right >= len(s) {
		return errors.New("Index is out of range.")
	}
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return nil
}
