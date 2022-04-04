package main

import "fmt"

func sortArray(nums []int) []int {
	var quicksort func(nums []int, left, right int) []int
	quicksort = func(nums []int, left, right int) []int {
		l, r := left, right
		if left > right {
			return nil
		}
		pivot := nums[left]
		for l < r {
			for l < r && pivot <= nums[r] {
				r--
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
			for l < r && pivot > nums[l] {
				l++
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
		}
		nums[l] = pivot
		quicksort(nums, left, l-1)
		quicksort(nums, l+1, right)
		return nums
	}
	return quicksort(nums, 0, len(nums)-1)
}

func main() {
	fmt.Println(sortArray([]int{4, 7, 9, 2, 4, 6, 7, 1}))
}
