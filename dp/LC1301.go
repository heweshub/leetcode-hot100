package main

import (
	"fmt"
	"math"
)

const inf = math.MinInt32

const mod1 = 1e9 + 7

var n int

func pathsWithMaxScore(board []string) []int {
	n = len(board)
	cs := make([][]byte, n)
	for i := range cs {
		cs[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			cs[i][j] = board[i][j]
		}
	}
	// 分别记录最大得分和最大方案数
	f := make([]int, n*n)
	g := make([]int, n*n)
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			idx := getIdx(i, j)
			if i == n-1 && j == n-1 {
				g[idx] = 1
				continue
			}
			if cs[i][j] == 'X' {
				f[idx] = inf
				continue
			}
			var val int
			if i == 0 && j == 0 {
				val = 0
			} else {
				val = int(cs[i][j] - '0')
			}
			u, t := inf, 0
			if i+1 < n {
				cur := f[getIdx(i+1, j)] + val
				cnt := g[getIdx(i+1, j)]
				res := update(cur, cnt, u, t)
				u = res[0]
				t = res[1]
			}
			if j+1 < n {
				cur := f[getIdx(i, j+1)] + val
				cnt := g[getIdx(i, j+1)]
				res := update(cur, cnt, u, t)
				u = res[0]
				t = res[1]
			}
			if i+1 < n && j+1 < n {
				cur := f[getIdx(i+1, j+1)] + val
				cnt := g[getIdx(i+1, j+1)]
				res := update(cur, cnt, u, t)
				u = res[0]
				t = res[1]
			}
			if u < 0 {
				f[idx] = inf
			} else {
				f[idx] = u
			}
			g[idx] = t
		}
	}
	ans := make([]int, 2)
	if f[getIdx(0, 0)] == inf {
		ans[0] = 0
		ans[1] = 0
	} else {
		ans[0] = f[getIdx(0, 0)]
		ans[1] = g[getIdx(0, 0)]
	}
	return ans
}

/*
u 表示最大得分，t 表示最大方案数
*/
func update(cur int, cnt int, u int, t int) []int {
	ans := []int{u, t}
	if cur > u {
		ans[0] = cur
		ans[1] = cnt
	} else if cur == u && cur != inf {
		// 相等的得分需要累加方案数
		ans[1] += cnt
	}
	ans[1] %= mod1
	return ans
}

// 二维数组转换为一维数组
func getIdx(x, y int) int {
	return x*n + y
}

func main() {
	board := []string{"E12", "1X1", "21S"}
	fmt.Println(pathsWithMaxScore(board))
}
