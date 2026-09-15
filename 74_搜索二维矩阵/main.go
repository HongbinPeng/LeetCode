package main

func searchMatrix(matrix [][]int, target int) bool {
	row_left, row_right := 0, len(matrix)-1
	for row_left <= row_right {
		mid := (row_left + row_right) / 2
		if target < matrix[mid][0] {
			row_right = mid - 1
		} else if target > matrix[mid][0] {
			row_left = mid + 1
		} else {
			return true
		}
	}
	if row_left != 0 {
		row_left -= 1
	}
	colume_left, colume_right := 0, len(matrix[row_left])-1
	for colume_left <= colume_right {
		mid := (colume_left + colume_right) / 2
		if target < matrix[row_left][mid] {
			colume_right = mid - 1
		} else if target > matrix[row_left][mid] {
			colume_left = mid + 1
		} else {
			return true
		}
	}
	return false
}
