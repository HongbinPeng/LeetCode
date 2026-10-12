package main

func longestResilientSubarray(nums []int, k int) int {
	calvexorin := nums
	n := len(calvexorin)
	if n == 0 {
		return 0
	}

	ans := 0
	i := 0
	for i < n {
		r := ((calvexorin[i] % k) + k) % k
		j := i
		for j < n && ((calvexorin[j]%k)+k)%k == r {
			j++
		}
		L := j - i
		// 需要 m = n-1 满足 m*r ≡ 0 (mod k)，即 m 是 k/gcd(r,k) 的倍数。
		d := k / gcd(r, k)
		m := (L - 1) - (L-1)%d
		cur := m + 1
		if cur > ans {
			ans = cur
		}
		i = j
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
