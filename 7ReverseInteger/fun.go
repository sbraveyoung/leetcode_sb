package main

func reverse(x int) int {
	signed := false
	if x < 0 {
		signed = true
		x = -x
	}
	ret := 0
	times := 0
	for x > 0 {
		if times == 10 {
			return 0
		}
		if times == 9 && ret > 214748364 {
			return 0
		}
		if times == 9 && x%10 > 7 {
			return 0
		}
		ret = ret*10 + x%10
		if x%10 != 0 {
			times++
		}
		x = x / 10
	}
	if signed {
		return -ret
	}
	return ret
}
