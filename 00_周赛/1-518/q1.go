package main

func countRotations(s string, k int) int {
	str := []rune(s)
	sumcount := 0
	ans := 0
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			ans++
		}
	}
	if ans == k {
		sumcount += 1
	}
	temp := str[len(str)-1]
	for j := 1; j < len(str); j++ {
		if str[j] == str[j-1] {
			ans--
		}
		if str[j-1] == temp {
			ans++
		}
		temp = str[j-1]
		if ans == k {
			sumcount += 1
		}

	}
	return sumcount
}
