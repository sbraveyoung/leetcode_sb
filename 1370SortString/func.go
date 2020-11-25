package sortString

import "sort"

type slice []rune

func (s slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s slice) Len() int           { return len(s) }
func (s slice) Less(i, j int) bool { return s[i] < s[j] }

func sortString(s string) string {
	dict := make(map[rune]int)
	for _, c := range s {
		dict[c]++
	}

	arr := slice{}
	for k := range dict {
		arr = append(arr, k)
	}
	sort.Sort(arr)

	var res []rune
	for len(dict) != 0 {
		for i := 0; i < len(arr); i++ {
			if c, ok := dict[arr[i]]; !ok {
				continue
			} else if c > 0 {
				res = append(res, arr[i])
				dict[arr[i]]--
			} else {
				delete(dict, arr[i])
			}
		}
		for i := len(arr) - 1; i >= 0; i-- {
			if c, ok := dict[arr[i]]; !ok {
				continue
			} else if c > 0 {
				res = append(res, arr[i])
				dict[arr[i]]--
			} else {
				delete(dict, arr[i])
			}
		}
	}
	return string(res)
}
