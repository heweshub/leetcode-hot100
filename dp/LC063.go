package main

func uniquePathsWithObstacles(grid [][]int) int {
	res := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		res[i] = make([]int, len(grid[0]))
	}
	if grid[0][0] == 0 {
		res[0][0] = 1
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				if i > 0 && j > 0 {
					res[i][j] = res[i][j-1] + res[i-1][j]
				} else if i > 0 {
					res[i][j] = res[i-1][j]
				} else if j > 0 {
					res[i][j] = res[i][j-1]
				}
			}
		}
	}
	return res[len(grid)-1][len(grid[0])-1]
}
