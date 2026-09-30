package main

func maxDepthAfterSplit(seq string) []int {
	str := []rune(seq)
	ans := make([]int, len(str))
	a, b := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			if a <= b {
				a++
				ans[i] = 0
			} else {
				b++
				ans[i] = 1
			}
		} else {
			if a >= b {
				a--
				ans[i] = 0
			} else {
				ans[i] = 1
				b--
			}
		}
	}
	return ans
}
