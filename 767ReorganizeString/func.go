package reorganizeString

import (
	"sort"
)

type Dict struct {
	char  int
	count int
}

type Slice []Dict

func (s Slice) Less(i, j int) bool {
	return s[i].count > s[j].count
}
func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Slice) Len() int {
	return len(s)
}

func reorganizeString(S string) string {
	if S == "" {
		return S
	}

	var array Slice
	for i := 0; i < 26; i++ {
		array = append(array, Dict{
			char:  'a' + i,
			count: 0,
		})
	}

	for _, c := range S {
		array[c-'a'].count++
	}
	sort.Sort(array)

	if array[0].count > len(S)-array[0].count+1 {
		return ""
	}

	res := make([]byte, len(S), len(S))
	index := 0
	for i := 0; len(array) != 0; i %= len(array) {
		// fmt.Printf("%+v\n", array)
		oldi := i
		if array[i].count != 0 {
			res[index] = byte(array[i].char)
			index++
			if i >= 1 {
				maxIndex := i - 1
				if i+1 < len(array) && array[i+1].count > array[maxIndex].count {
					maxIndex = i + 1
				}
				i = maxIndex
			} else { //i==0
				i++
			}
			array[oldi].count--
		}
		if array[oldi].count == 0 {
			array = append(array[:oldi], array[oldi+1:]...)
			i = 0
		}
		if len(array) == 0 {
			break
		}
	}
	return string(res)
}
