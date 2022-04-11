package main

import (
	"fmt"
	"math"
)

func minFallingPathSum2(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = grid[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt64
			var val = grid[i][j]
			for k := 0; k < n; k++ {
				if k != j {
					dp[i][j] = min(dp[i][j], dp[i-1][k]+val)
				}
			}
		}
	}
	fmt.Println(dp)
	ans := math.MaxInt64
	for i := 0; i < n; i++ {
		ans = min(ans, dp[n-1][i])
	}
	return ans
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(minFallingPathSum2([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
