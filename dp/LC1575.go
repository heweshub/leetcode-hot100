package main

import (
	"fmt"
	"math"
)

func countRoutes(locations []int, start int, finish int, fuel int) int {
	mod := 1000000007
	n := len(locations)
	// 初始化当前下标下的油量
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, fuel+1)
		for j := 0; j < fuel+1; j++ {
			cache[i][j] = -1
		}
	}
	// 构造dfs
	var dfs func(loc []int, cur int, end int, fuel int) int
	dfs = func(loc []int, cur int, end int, fuel int) int {
		// 值不为-1，表示已经计算过直接返回即可
		if cache[cur][fuel] != -1 {
			return cache[cur][fuel]
		}
		// 当前点到终点的fuel油量
		need := int(math.Abs(float64(loc[cur] - loc[end])))
		// 油不够的话直接将当前下标下的油量置为0，表示不能到达
		if need > fuel {
			cache[cur][fuel] = 0
			return 0
		}
		n := len(loc)
		sum := 0
		// 相等也是一种方案
		if cur == end {
			sum = 1
		}
		for i := 0; i < n; i++ {
			if i != cur {
				need := int(math.Abs(float64(loc[i] - loc[cur])))
				// 油够用的情况，继续dfs
				if fuel >= need {
					sum += dfs(loc, i, end, fuel-need)
					sum %= mod
				}
			}
		}
		cache[cur][fuel] = sum
		return sum
	}
	ans := dfs(locations, start, finish, fuel)
	return ans
}

func main() {

	fmt.Println(countRoutes([]int{4, 3, 1}, 1, 0, 6))
}
