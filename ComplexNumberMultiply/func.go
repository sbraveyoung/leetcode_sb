package complex

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	c, d, x, y := do(a, b)
	return fmt.Sprintf("%d+%di", c*x-d*y, c*y+d*x)
}

func do(a, b string) (int, int, int, int) {
	two := strings.Split(a, "+")
	c, err0 := strconv.Atoi(two[0])
	d, err1 := strconv.Atoi(strings.TrimSuffix(two[1], "i"))
	two = strings.Split(b, "+")
	x, err2 := strconv.Atoi(two[0])
	y, err3 := strconv.Atoi(strings.TrimSuffix(two[1], "i"))

	if err0 != nil || err1 != nil || err2 != nil || err3 != nil {
		panic("input error")
	}
	return c, d, x, y
}
