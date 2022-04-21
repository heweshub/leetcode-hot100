package main

import (
	"fmt"
)

func maxProfit3(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

func main() {
	fmt.Println(maxProfit3([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
}