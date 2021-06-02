package modifyString

func modifyString(s string) string {
	if s == "" {
		return s
	}

	b := []rune(s)
	for index, _ := range b {
		if b[index] != '?' {
			continue
		}

		for _, c := range "abcdefghijklmnopqrstuvwxyz" {
			if index != 0 && c == b[index-1] {
				continue
			}
			if index != len(b)-1 && c == b[index+1] {
				continue
			}
			b[index] = c
			break
		}
	}
	return string(b)
}
