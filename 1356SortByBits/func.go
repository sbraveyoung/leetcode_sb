package sortByBits

import "sort"

type slice []int

func (s slice) Len() int      { return len(s) }
func (s slice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s slice) Less(i, j int) bool {
	bitsI := getBits(s[i])
	bitsJ := getBits(s[j])
	if bitsI < bitsJ {
		return true
	} else if bitsI == bitsJ {
		return s[i] < s[j]
	}
	return false
}

func sortByBits(arr []int) []int {
	s := slice(arr)
	sort.Sort(s)
	return []int(s)
}

func getBits(a int) int {
	bits := 0
	for a != 0 {
		a &= (a - 1)
		bits++
	}
	return bits
}
