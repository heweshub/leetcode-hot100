package main

import "math"

var MAX = math.MaxInt32

func minFallingPathSum(mat [][]int) (ans int) {
	ans = MAX
	for i := 0; i < len(mat); i++ {
		ans = min(ans, find(mat, i))
	}
	return
}

func find(mat [][]int, u int) (ans int) {
	n := len(mat)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if i == u {
			f[0][i] = mat[0][i]
		} else {
			f[0][i] = MAX
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			f[i][j] = MAX
			val := mat[i][j]
			if f[i-1][j] != MAX {
				f[i][j] = min(f[i][j], f[i-1][j]+val)
			}
			if j-1 >= 0 && f[i-1][j-1] != MAX {
				f[i][j] = min(f[i][j], f[i-1][j-1]+val)
			}
			if j+1 < n && f[i-1][j+1] != MAX {
				f[i][j] = min(f[i][j], f[i-1][j+1]+val)
			}
		}
	}
	ans = MAX
	for p := 0; p < n; p++ {
		ans = min(ans, f[n-1][p])
	}
	return
}
