package main

func uniquePaths(m int, n int) int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	res[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				res[i][j] = res[i][j-1] + res[i-1][j]
			} else if i > 0 {
				res[i][j] = res[i-1][j]
			} else if j > 0 {
				res[i][j] = res[i][j-1]
			}
		}
	}
	return res[m-1][n-1]
}
