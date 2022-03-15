package main

import "fmt"

// 11. 盛最多水的容器
// 问题描述：给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
// 找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	// 分别记录容器的最大体积和当前体积
	maxSize := 0
	size := 0
	for l < r {
		// 左右指针的坐标和对应数组中的较小值求得当前容器的体积
		size = (r - l) * min(height[l], height[r])
		if size > maxSize {
			maxSize = size
		}

		// 数组中对应值较低的指针需要移动，才能周到更大的体积
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return maxSize
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
