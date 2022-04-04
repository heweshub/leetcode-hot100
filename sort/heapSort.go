package main

import "fmt"

func sortArray2(nums []int) []int {
	var heapify func(nums []int, root, end int)
	// 调整大根堆
	heapify = func(nums []int, root, end int) {
		for {
			child := root*2 + 1
			if child > end {
				return
			}
			if child < end && nums[child] <= nums[child+1] {
				child++
			}
			if nums[root] > nums[child] {
				return
			}
			// 不满足上述条件的话就要调整root和child的位置
			nums[root], nums[child] = nums[child], nums[root]
			root = child
		}
	}
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		heapify(nums, i, end)
	}
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		heapify(nums, 0, end)
	}
	return nums
}

func main() {
	fmt.Println(sortArray2([]int{4, 7, 9, 2, 4, 6, 7, 1}))
}
