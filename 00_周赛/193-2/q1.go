package main

import "sort"

func maxProductPair(nums []int, target int) []int {
	// 记录每个值最后一次出现的下标。两个不同的值必然对应不同下标。
	index := make(map[int]int)
	for i, v := range nums {
		index[v] = i
	}

	// 收集去重后的值并排序。
	vals := make([]int, 0, len(index))
	for v := range index {
		vals = append(vals, v)
	}
	sort.Ints(vals)

	ans := []int{-1, -1}
	best := 0
	found := false

	l, r := 0, len(vals)-1
	for l < r {
		sum := vals[l] + vals[r]
		if sum == target {
			prod := vals[l] * vals[r]
			if !found || prod > best {
				best = prod
				// vals[r] 更大，作为 nums[i]；vals[l] 更小，作为 nums[j]。
				ans = []int{index[vals[r]], index[vals[l]]}
				found = true
			}
			l++
			r--
		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	return ans
}
