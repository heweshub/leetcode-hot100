package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	l, r := 0, n-1
	var find func(nums []int, l, r, target int) int
	find = func(nums []int, l, r, target int) int {
		for l < r {
			mid := (l + r) / 2
			if nums[mid] == target {
				return mid
			}
			if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return -1
	}
	idx := find(nums, l, r, target)
	if idx == -1 {
		return []int{-1, -1}
	}
	l, r = idx, idx
	for l > 1 {
		if nums[l-1] != target {
			break
		}
		l--
	}
	for r < n-2 {
		if nums[r+1] != target {
			break
		}
		r++
	}
	return []int{l, r}
}

func main() {
	//输入：nums = [5,7,7,8,8,10], target = 8
	//输出：[3,4]
	fmt.Println(searchRange([]int{1}, 1))
}
