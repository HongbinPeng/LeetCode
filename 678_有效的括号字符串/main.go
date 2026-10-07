package main

func checkValidString(s string) bool {
	lo, hi := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lo++
			hi++
		} else if ch == ')' {
			if lo > 0 {
				lo--
			}
			hi--
			if hi < 0 { 
				return false
			}
		} else {
			if lo > 0 {
				lo--
			}
			hi++
		}
	}
	return lo == 0
}
