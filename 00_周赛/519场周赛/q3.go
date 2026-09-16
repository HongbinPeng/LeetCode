package main

func shadowPairs(nums []int) int64 {
	type group struct {
		value int
		count int
	}
	stack := make([]group, 0, len(nums))
	activeCount := 0
	var ans int64
	for _, x := range nums {
		for len(stack) > 0 && stack[len(stack)-1].value > x {
			activeCount -= stack[len(stack)-1].count
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 && stack[len(stack)-1].value == x {
			ans += int64(activeCount - stack[len(stack)-1].count)
			stack[len(stack)-1].count++
		} else {
			ans += int64(activeCount)
			stack = append(stack, group{value: x, count: 1})
		}
		activeCount++
	}
	return ans
}
