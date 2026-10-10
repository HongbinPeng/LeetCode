package main

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
	d := make([]int, n)
	hi := 0
	for i := 0; i < n; i++ {
		diff := nums1[i] - nums2[i]
		if diff < 0 {
			diff = -diff
		}
		d[i] = diff
		if diff > hi {
			hi = diff
		}
	}
	k := k1 + k2
	lo, hi2 := 0, hi
	for lo < hi2 {
		mid := (lo + hi2) / 2
		var cost int64
		for _, v := range d {
			if v > mid {
				cost += int64(v - mid)
			}
		}
		if cost <= int64(k) {
			hi2 = mid
		} else {
			lo = mid + 1
		}
	}
	x := lo
	r := int64(k)
	for _, v := range d {
		if v > x {
			r -= int64(v - x)
		}
	}
	var ans int64
	for _, v := range d {
		if v < x {
			ans += int64(v) * int64(v)
		} else {
			val := int64(x)
			if r > 0 {
				val--
				r--
			}
			if val < 0 {
				val = 0
			}
			ans += val * val
		}
	}
	return ans
}
