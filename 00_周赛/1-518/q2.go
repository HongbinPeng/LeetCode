package main

func countGoodRotations(nums []int) int {
	slow, quick := 0, len(nums)/2
	sumsub := 0
	dp := make([]int, len(nums))
	for quick < len(nums) {
		temp := nums[slow] - nums[quick]
		dp[slow] = temp
		sumsub += temp
		slow++
		quick++
	}
	for i := slow; i < len(nums); i++ {
		dp[i] = -dp[i-slow]
	}
	ans := 0
	if sumsub > 0 {
		ans += 1
	}
	for i := 1; i < len(nums); i++ {
		sumsub -= 2 * dp[i-1]
		if sumsub > 0 {
			ans++
		}
	}
	return ans
}
