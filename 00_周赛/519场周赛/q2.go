package main

func minOperations(nums []int) int64 {
	var ans int64

	for _, num := range nums {
		x := int64(num)
		length := digitCount(x)
		best := int64(1 << 62)

		// 最近的回文数只可能有 length-1、length、length+1 位。
		left := length - 1
		if left < 1 {
			left = 1
		}

		for palindromeLength := left; palindromeLength <= length+1; palindromeLength++ {
			halfLength := (palindromeLength + 1) / 2
			unit := power10(halfLength - 1)

			var targetPrefix int64

			switch {
			case palindromeLength < length:
				targetPrefix = unit*10 - 1
			case palindromeLength > length:
				targetPrefix = unit
			default:
				targetPrefix = x / power10(palindromeLength-halfLength)
			}

			// 回文数的奇偶性由首位（同时也是末位）决定。
			firstDigit := int64(1)
			if x%2 == 0 {
				firstDigit = 2
			}

			for ; firstDigit <= 9; firstDigit += 2 {
				low := firstDigit * unit
				high := (firstDigit+1)*unit - 1

				prefix := targetPrefix
				if prefix < low {
					prefix = low
				}
				if prefix > high {
					prefix = high
				}

				// 镜像值关于 prefix 单调，只需检查附近三个前缀。
				for delta := int64(-1); delta <= 1; delta++ {
					currentPrefix := prefix + delta
					if currentPrefix < low || currentPrefix > high {
						continue
					}

					palindrome := makePalindrome(
						currentPrefix,
						palindromeLength,
					)

					difference := x - palindrome
					if difference < 0 {
						difference = -difference
					}

					if difference < best {
						best = difference
					}
				}
			}
		}

		ans += best / 2
	}

	return ans
}

func makePalindrome(prefix int64, length int) int64 {
	result := prefix

	if length%2 == 1 {
		prefix /= 10
	}

	for prefix > 0 {
		result = result*10 + prefix%10
		prefix /= 10
	}

	return result
}

func digitCount(x int64) int {
	count := 0
	for x > 0 {
		count++
		x /= 10
	}
	return count
}

func power10(exponent int) int64 {
	result := int64(1)
	for i := 0; i < exponent; i++ {
		result *= 10
	}
	return result
}
