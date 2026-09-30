package main

func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if (m+n-1)%2 != 0 || grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}
	// visited[i][j][k]: 到 (i,j) 且 balance=k 是否访问过
	visited := make([][][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			visited[i][j] = make([]bool, m+n) // balance 范围 0 ~ m+n
		}
	}

	var dfs func(i, j, balance int) bool
	dfs = func(i, j, balance int) bool {
		if i >= m || j >= n || balance < 0 {
			return false
		}
		if visited[i][j][balance] {
			return false
		}
		visited[i][j][balance] = true

		if grid[i][j] == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
		if i == m-1 && j == n-1 {
			return balance == 0
		}
		return dfs(i+1, j, balance) || dfs(i, j+1, balance)
	}

	return dfs(0, 0, 0)
}
