package main

func resultArray(nums []int, k int) []int64 {
	ans := make([]int64, k)
	dp := make([][]int64, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int64, k)

	}
	for i := 0; i < len(nums); i++ {
		dp[i+1][nums[i]%k] = 1
		for r := 0; r < k; r++ {
			dp[i+1][(r*nums[i])%k] = dp[i][r] + dp[i+1][(r*nums[i])%k]
		}
		for x, c := range dp[i+1] {
			ans[x] += c
		}
	}
	return ans
}
