package main

func minDays(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		// 可行方案：每次只得1分，中间各跳过一天。
		dp[i] = 2*i - 1

		// 枚举最后一段连续得分的天数。
		for k := 1; k*(k+1)/2 <= i; k++ {
			points := k * (k + 1) / 2
			days := k

			if points < i {
				// 前面凑出剩余分数，再跳过一天，接上最后一段。
				days += dp[i-points] + 1
			}

			if days < dp[i] {
				dp[i] = days
			}
		}
	}

	return dp[n]
}
