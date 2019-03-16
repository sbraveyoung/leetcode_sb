package generateParenthesis

import "fmt"

func generateParenthesis(n int) []string {
	res := []string{}
	if n <= 0 {
		return res
	}
	generate("", 0, 0, n, &res)
	return res
}

func generate(s string, left, right, n int, res *[]string) {
	if right >= n {
		*res = append(*res, s)
		return
	}
	if left >= n {
		for i := 0; i < n-right; i++ {
			s = fmt.Sprintf("%s)", s)
		}
		*res = append(*res, s)
		return
	}
	if left > right {
		generate(fmt.Sprintf("%s(", s), left+1, right, n, res)
		generate(fmt.Sprintf("%s)", s), left, right+1, n, res)
	} else {
		generate(fmt.Sprintf("%s(", s), left+1, right, n, res)
	}
}

//too slow
// func generateParenthesis(n int) []string {
// res := []string{}
// if n <= 0 {
// return res
// }
// generate("", 0, 0, n, &res)
// return res
// }

// func generate(s string, left, right, n int, res *[]string) {
// if right >= n {
// *res = append(*res, s)
// return
// }
// if left >= n {
// for i := 0; i < n-right; i++ {
// s += ")"
// }
// *res = append(*res, s)
// return
// }
// if left > right {
// generate(s+"(", left+1, right, n, res)
// generate(s+")", left, right+1, n, res)
// } else {
// generate(s+"(", left+1, right, n, res)
// }
// }
