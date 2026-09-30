package main

func minSumOfLengths(arr []int, target int) int {
	pos := make(map[int]int, 0)
	s := 0
	pos[0] = -1
	ans, minL := len(arr)+1, len(arr)

	for i, num := range arr {
		s += num
		if j, ok := pos[s-target]; ok {
			length := i - j
			prev := len(arr)
			if j != -1 {
				prev = arr[j]
			}
			ans = min(ans, prev+length)
			if length < minL {
				minL = length
			}
		}
		arr[i] = minL
		pos[s] = i
	}
	if ans == len(arr)+1 {
		return -1
	}
	return ans
}
