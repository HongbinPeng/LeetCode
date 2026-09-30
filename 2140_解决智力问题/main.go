package main

func mostPoints(questions [][]int) int64 {
	dp := make([]int64, len(questions))
	for i := len(questions) - 1; i >= 0; i-- {
		if i == len(questions)-1 {
			dp[i] = int64(questions[i][0])
		} else {
			temp := i + questions[i][1] + 1
			if temp < len(questions) {
				dp[i] = max(dp[temp]+int64(questions[i][0]), dp[i+1])
			} else {
				dp[i] = max(int64(questions[i][0]), dp[i+1])
			}
		}
	}
	return dp[0]
}
