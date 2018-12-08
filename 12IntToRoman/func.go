package intToRoman

import "fmt"

func intToRoman(num int) string {
	baseTable := [...][2]interface{}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var romanString string
	for num != 0 {
		for i := 0; i < len(baseTable); i++ {
			if num >= baseTable[i][0].(int) {
				romanString = fmt.Sprintf("%s%s", romanString, baseTable[i][1].(string))
				num -= baseTable[i][0].(int)
				break
			}
		}
	}
	return romanString
}
