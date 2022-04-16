package main

func maxProfit(prices []int) int {
	maxPro := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			maxPro += prices[i+1] - prices[i]
		}
	}
	return maxPro
}

func main() {
	maxProfit([]int{1, 2, 3, 4, 5})
}
