package main

import "fmt"

const mod int = 1e9 + 7

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findPath(m, n, maxMove, startRow, startColumn int) (ans int) {
	dp := make([][][]int, maxMove+1)
	// 构造dp矩阵
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	// 原始位置的步数肯定为1
	dp[0][startRow][startColumn] = 1
	// 相当于从startRow、startColumn这个点以当前最大步数向外扩展直到出界，计算总的count数
	for i := 0; i < maxMove; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				count := dp[i][j][k]
				if count > 0 {
					// 计算i+1次移动的情况
					for _, dir := range dirs {
						j1, k1 := j+dir.x, k+dir.y
						// 没出界的情况，更新dp上下左右的值
						if j1 >= 0 && j1 < m && k1 >= 0 && k1 < n {
							dp[i+1][j1][k1] = (dp[i+1][j1][k1] + count) % mod
						} else { // 出界就加入到总的次数中
							ans = (ans + count) % mod
						}
					}
				}
				fmt.Println(dp)
			}
		}
	}
	//fmt.Println(dp)
	return
}

func main() {
	m := 1
	n := 3
	maxMove := 3
	startRow := 0
	startColumn := 1
	fmt.Println(findPath(m, n, maxMove, startRow, startColumn))
}
