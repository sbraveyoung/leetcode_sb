package titleToNumber

//https://leetcode-cn.com/problems/excel-sheet-column-number/

func titleToNumber(s string) int {
	num := 0
	for _, c := range s {
		num = num*26 + int(c-'A'+1)
	}
	return num
}
