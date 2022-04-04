package main

import "fmt"

func sortArray1(nums []int) []int {
	merge := func(left, right []int) []int {
		res := make([]int, len(left)+len(right))
		var r, l, i int
		for l < len(left) && r < len(right) {
			if left[l] < right[r] {
				res[i] = left[l]
				l++
			} else {
				res[i] = right[r]
				r++
			}
			i++
		}
		// 拷贝之后剩余的元素
		copy(res[i:], left[l:])
		copy(res[i:], right[r:])
		return res
	}
	var sort func(nums []int) []int
	sort = func(nums []int) []int {
		if len(nums) <= 1 {
			return nums
		}
		mid := len(nums) / 2
		left := sort(nums[:mid])
		right := sort(nums[mid:])
		return merge(left, right)
	}
	return sort(nums)
}

func main() {
	fmt.Println(sortArray1([]int{4, 7, 9, 2, 4, 6, 7, 1}))
}
