package restoreipaddresses

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	addresses := []string{}
	if len(s) < 4 || len(s) > 12 {
		return addresses
	}
	split(s, "", "", 1, &addresses)
	return addresses
}

func split(s, ip, curPart string, partNum int, addresses *[]string) {
	if len(s) == 0 {
		return
	}
	if partNum == 4 {
		if isPartofIP(s) {
			*addresses = append(*addresses, ip+s)
		}
		return
	}
	switch len(curPart) {
	case 0:
		split(s[1:], ip, curPart+s[0:1], partNum, addresses)
	case 3:
		if isPartofIP(curPart) {
			split(s, ip+curPart+".", "", partNum+1, addresses)
		}
	default: //1,2
		split(s[1:], ip, curPart+s[0:1], partNum, addresses)
		if isPartofIP(curPart) {
			split(s, ip+curPart+".", "", partNum+1, addresses)

		}
	}
}

func isPartofIP(s string) (res bool) {
	if s[0] == '0' && len(s) != 1 {
		return false
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		res = false
		return res
	}
	if n > 0 && n <= 255 {
		res = true
		return res
	}
	if n == 0 && s == "0" {
		res = true
		return res
	}
	res = false
	return res
}
