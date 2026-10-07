package main

func removeInvalidParentheses(s string) (ans []string) {
	target, left := 0, 0
	for _, ch := range s {
		if ch == '(' {
			target++
			left++
		} else if ch == ')' && left > 0 {
			left--
		}
	}
	target = target - left
	path := []byte{}
	ans = []string{}
	n := len(s)
	var dfs func(i, left, right int)
	dfs = func(i, left, right int) {
		if left < right || left > target || left+right+n-i < 2*target {
			return
		}
		if i == n {
			ans = append(ans, string(path))
			return
		}
		ch := s[i]
		if ch == '(' || ch == ')' {
			j := i + 1
			for j < n && s[j] == ch {
				j++
			}
			dfs(j, left, right)
		}
		if ch == '(' {
			left++
		} else if ch == ')' {
			right++
		}
		path = append(path, ch)
		dfs(i+1, left, right)
		path = path[0 : len(path)-1]
	}
	dfs(0, 0, 0)
	return ans
}
