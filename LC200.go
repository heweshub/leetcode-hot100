package main

// LC200岛屿数量
func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	// 额外构造辅助矩阵记录已访问的情况
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
			return
		}
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 在判断新岛屿的条件上首先是grid[i][j]=1，还有就是为访问过的点
			if grid[i][j] == '1' && visited[i][j] == false {
				ans++
				dfs(i, j)
			}
		}
	}
	return
}
