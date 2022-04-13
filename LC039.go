package main

import "fmt"

//func combinationSum(c []int, target int) (ans [][]int) {
//	comb := []int{}
//	var dfs func(target, idx int)
//	dfs = func(target, idx int) {
//		if idx == len(c) {
//			return
//		}
//		if target == 0 {
//			ans = append(ans, append([]int{}, comb...))
//			return
//		}
//
//		dfs(target, idx+1)
//		if target-c[idx] >= 0 {
//			comb = append(comb, c[idx])
//			dfs(target-c[idx], idx)
//			comb = comb[:len(comb)-1]
//		}
//	}
//	dfs(target, 0)
//	return
//}

func combinationSum(candidates []int, target int) (ans [][]int) {
	var dfs func(start int, tmp []int, sum int)
	dfs = func(start int, tmp []int, sum int) {
		if sum >= target {
			// 满足target的组合加入到ans中
			if sum == target {
				newTmp := make([]int, len(tmp))
				// 切片的复制，是地址的引用
				copy(newTmp, tmp)
				ans = append(ans, newTmp)
			}
			return
		}
		// 从start开始的话可以有效的避免组合重复问题，由题意可知，candidates中的元素都是不重复的
		for i := start; i < len(candidates); i++ {
			// 将candidates[i]加入到tmp中进行dfs
			tmp = append(tmp, candidates[i])
			dfs(i, tmp, sum+candidates[i])
			// 把candidates[i]剔除掉继续下一轮循环
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return
}

func main() {
	candidates := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(candidates, 7))
}
