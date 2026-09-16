package main

func searchRange(nums []int, target int) []int {
	ans := make([]int, 2)
	ans[0] = -1
	ans[1] = -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			ans[0] = mid
			right = mid - 1
		}
	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			ans[1] = mid
			left = mid + 1
		}
	}
	return ans
}
