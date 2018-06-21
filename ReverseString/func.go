package reverseString

import "unsafe"

func reverseString(s string) string {
	length := len(s)
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		ret[i] = s[length-i-1]
	}
	return bytes2str(ret)
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
