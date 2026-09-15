package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = left + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}
