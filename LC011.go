package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxA := 0
	for l < r {
		cur := min(height[l], height[r]) * (r - l)
		if cur > maxA {
			maxA = cur
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxA
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
