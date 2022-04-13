package main

func minPathSum(grid [][]int) int {
	ans := grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i > 0 && j > 0 {
				if grid[i][j-1] > grid[i-1][j] {
					ans[i][j] = grid[i][j] + grid[i-1][j]
				} else {
					ans[i][j] = grid[i][j] + grid[i][j-1]
				}
			} else if i > 0 {
				ans[i][j] = grid[i][j] + ans[i-1][j]
			} else if j > 0 {
				ans[i][j] = grid[i][j] + ans[i][j-1]
			}
		}
	}
	return ans[len(grid)-1][len(grid[0])-1]
}
