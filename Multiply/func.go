package multiply

import (
	"fmt"
	"strings"
)

func multiply(num1, num2 string) (mult string) {
	baseNum := map[byte]int{
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'0': 0,
	}
	length1, length2 := len(num1), len(num2)
	result := make([][]int, length1)
	slot := 0
	for i := length1 - 1; i >= 0; i-- {
		result[i] = make([]int, length2+1)
		for j := length2 - 1; j >= 0; j-- {
			result[i][j+1] = baseNum[num1[i]]*baseNum[num2[j]] + slot
			if result[i][j+1]/10 > 0 {
				slot = result[i][j+1] / 10
				result[i][j+1] %= 10
			} else {
				slot = 0
			}
		}
		if slot != 0 {
			result[i][0] = slot
		}
		slot = 0
	}

	slot1 := 0
	for k := length1 - 1; k >= 0; k-- {
		sum := slot1
		slot1 = 0
		for i, j := k, length2; i < length1 && j >= 0; i, j = i+1, j-1 {
			sum += result[i][j]
		}
		if sum/10 > 0 {
			slot1 = sum / 10
			sum %= 10
		}
		mult = fmt.Sprintf("%d%s", sum, mult)
	}
	for k := length2 - 1; k >= 0; k-- {
		sum := slot1
		slot1 = 0
		for i, j := 0, k; i < length1 && j >= 0; i, j = i+1, j-1 {
			sum += result[i][j]
		}
		if sum/10 > 0 {
			slot1 = sum / 10
			sum %= 10
		}
		mult = fmt.Sprintf("%d%s", sum, mult)
	}
	mult = TrimPrefixZero(mult)
	return mult
}

func TrimPrefixZero(str string) string {
	for len(str) > 1 {
		if str[0] == '0' {
			str = strings.TrimPrefix(str, "0")
		} else {
			break
		}
	}
	return str
}
