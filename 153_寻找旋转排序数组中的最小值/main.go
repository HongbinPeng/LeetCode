package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	ans := nums[0]
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= nums[left] {
			ans = min(ans, nums[left])
			left = mid + 1
		} else {
			ans = min(ans, nums[mid])
			right = mid - 1
		}
	}
	return ans
}
