package main

type flag struct {
	count int
	dis   int
	last  int
	plus  bool
}

func countSpecialIntegers(nums []int) int {
	a := make(map[int]*flag, 0)
	ans := 0
	for indx, num := range nums {
		if mesg, ok := a[num]; ok {
			if mesg.count == 1 {
				mesg.count++
				mesg.dis = indx - mesg.last
				mesg.last = indx
				continue
			} else if mesg.count == 2 {
				if indx-mesg.last == mesg.dis {
					mesg.count++
					mesg.last = indx
					mesg.plus = true
					ans++
				} else {
					mesg.count = 3
					mesg.dis = -1
					mesg.last = indx
				}
			} else {
				if mesg.plus {
					mesg.plus = false
					mesg.count = 3
					ans--
				}
			}
		} else {
			a[num] = &flag{count: 1, last: indx}
		}
	}
	return ans
}
