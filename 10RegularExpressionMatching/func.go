package RegularExpressionMatching

import "container/list"

type args struct {
	s     string
	p     string
	pprev byte
}

func isMatch(s string, p string) bool {
	var stack list.List
	return match(s, p, 0, stack)
}

func match(s string, p string, pprev byte, stack list.List) bool {
	stack.PushBack(args{s: s, p: p, pprev: pprev})
	for stack.Len() > 0 {
		arg, _ := stack.Back().Value.(args)
		stack.Remove(stack.Back())
		s = arg.s
		p = arg.p
		pprev = arg.pprev
		if len(s) == 0 && len(p) == 0 {
			return true
		}
		if len(p) == 0 {
			continue
		}
		if len(s) == 0 {
			switch p[0] {
			case '.':
				if 1 < len(p) && p[1] == '*' {
					stack.PushBack(args{s: s, p: p[2:], pprev: 0})
					continue
				}
				continue
			case '*':
				if pprev == 0 {
					continue
				} else {
					stack.PushBack(args{s: s, p: p[1:], pprev: 0})
					continue
				}
			default:
				if 1 < len(p) && p[1] == '*' {
					stack.PushBack(args{s: s, p: p[2:], pprev: 0})
					continue
				}
				continue
			}
		}

		switch p[0] {
		case '.':
			if 1 < len(p) && p[1] == '*' {
				stack.PushBack(args{s: s, p: p[2:], pprev: 0})
				stack.PushBack(args{s: s[1:], p: p, pprev: p[0]})
				continue
			}
			stack.PushBack(args{s: s[1:], p: p[1:], pprev: p[0]})
			continue
		case '*':
			if pprev == s[0] || pprev == '.' {
				stack.PushBack(args{s: s, p: p[2:], pprev: 0})
				stack.PushBack(args{s: s[1:], p: p, pprev: s[0]})
				continue
			} else {
				stack.PushBack(args{s: s, p: p[1:], pprev: 0})
				continue
			}
		default:
			if s[0] != p[0] {
				if 1 < len(p) && p[1] == '*' {
					stack.PushBack(args{s: s, p: p[2:], pprev: 0})
					continue
				}
				continue
			} else {
				if 1 < len(p) && p[1] == '*' {
					stack.PushBack(args{s: s, p: p[2:], pprev: 0})
					stack.PushBack(args{s: s[1:], p: p, pprev: s[0]})
					continue
				} else {
					stack.PushBack(args{s: s[1:], p: p[1:], pprev: s[0]})
					continue
				}
			}
		}
	}
	return false
}

//func isMatch(s string, p string) bool {
//	return match(s, p, 0, stack)
//}
//
//func match(s string, p string, pprev byte) bool {
//	for len(stack) > 0 {
//		if len(s) == 0 && len(p) == 0 {
//			return true
//		}
//		if len(p) == 0 {
//			return false
//		}
//		if len(s) == 0 {
//			switch p[0] {
//			case '.':
//				if 1 < len(p) && p[1] == '*' {
//					return match(s, p[2:], 0)
//				}
//				return false
//			case '*':
//				if pprev == 0 {
//					return false
//				} else {
//					return match(s, p[1:], 0)
//				}
//			default:
//				if 1 < len(p) && p[1] == '*' {
//					return match(s, p[2:], 0)
//				}
//				return false
//			}
//		}
//
//		switch p[0] {
//		case '.':
//			if 1 < len(p) && p[1] == '*' {
//				return match(s, p[2:], 0) || match(s[1:], p, p[0])
//			}
//			return match(s[1:], p[1:], p[0])
//		case '*':
//			if pprev == s[0] || pprev == '.' {
//				return match(s, p[1:], 0) || match(s[1:], p, s[0])
//			} else {
//				return match(s, p[1:], 0)
//			}
//		default:
//			if s[0] != p[0] {
//				if 1 < len(p) && p[1] == '*' {
//					return match(s, p[2:], 0)
//				}
//				return false
//			} else {
//				if 1 < len(p) && p[1] == '*' {
//					return match(s, p[2:], 0) || match(s[1:], p, s[0])
//				} else {
//					return match(s[1:], p[1:], s[0])
//				}
//			}
//		}
//	}
//}
