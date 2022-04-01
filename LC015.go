package main

func threeSum(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i && nums[j] == nums[j-1] {
				continue
			}
			r := n - 1
			sum := nums[i] + nums[j] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				continue
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[r]})
				for j < r && nums[r] == nums[r-1] {
					r--
				}
			}
		}
	}
	return ans
}
