package main

func removeOuterParentheses(s string) string {
	st := []rune{}
	dpth := 0
	for _, ch := range s {
		if ch == '(' {
			if dpth > 0 {
				st = append(st, ch)
			}
			dpth++
		} else {
			dpth--
			if dpth > 0 {
				st = append(st, ch)
			}
		}
	}
	return string(st)
}
