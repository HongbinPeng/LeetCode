package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	x11, x12 := rec1[0], rec1[2]
	y11, y12 := rec1[1], rec1[3]
	x21, x22 := rec2[0], rec2[2]
	y21, y22 := rec2[1], rec2[3]
	if (x21 > x11 && x21 < x12) || (x22 > x11 && x22 < x12) || (x11 > x21 && x11 < x22) || (x12 > x21 && x12 < x22) {
		if (y21 > y11 && y21 < y12) || (y22 > y11 && y22 < y12) || (y11 > y21 && y11 < y22) || (y12 > y21 && y12 < y22) {
			return true
		}
	}
	if x11 == x21 && x12 == x22 && y11 == y21 && y12 == y22 {
		return true
	}
	return false
}
