package main

func minInsertions(s string) int {
	ans := 0
	l := 0
	str := []rune(s)
	n := len(str)
	for i := 0; i < len(s); i++ {
		if str[i] == '(' {
			l++
		} else {
			if l > 0 {
				l--
				i += 1
				if i < n {
					if str[i] == ')' {
						continue
					} else {
						ans += 1
						l++
						continue
					}
				} else {
					ans += 1
					break
				}
			} else {
				ans += 1
				l++
				i -= 1
			}
		}
	}
	return ans + 2*l
}
