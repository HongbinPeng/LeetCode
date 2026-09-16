package main

func cyclicShift(n int, grid [][]int, rowShift []int, colShift []int) [][]int {
	for idx, val := range rowShift {
		move := val % n
		if move == 0 {
			continue
		}
		move = n - move
		reverse(grid[idx])
		reverse(grid[idx][:move])
		reverse(grid[idx][move:])
	}
	for idx, val := range colShift {
		move := val % n
		if move == 0 {
			continue
		}
		move = n - move
		reverse2(grid, false, idx, 0, len(grid)-1)
		reverse2(grid, false, idx, 0, move-1)
		reverse2(grid, false, idx, move, len(grid)-1)
	}
	return grid
}
func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
func reverse2(a [][]int, isrow bool, guding, left, right int) {
	if isrow {
		for i, j := left, right; i < j; i, j = i+1, j-1 {
			a[guding][i], a[guding][j] = a[guding][j], a[guding][i]
		}
	} else {
		for i, j := left, right; i < j; i, j = i+1, j-1 {
			a[i][guding], a[j][guding] = a[j][guding], a[i][guding]
		}
	}
}
