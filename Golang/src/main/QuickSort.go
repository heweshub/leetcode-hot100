package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func QuickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	left, right := l, r
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		if left < right {
			nums[left] = nums[right]
		}
		for left < right && nums[left] <= pivot {
			left++
		}
		if left < right {
			nums[right] = nums[left]
		}
		if left >= right {
			nums[left] = pivot
		}
	}
	QuickSort(nums, l, left-1)
	QuickSort(nums, left+1, r)
}

func main() {
	//var num int
	//fmt.Scanln(&num)
	//var nums []int
	//for i:=0;i<num;i++{
	//	fmt.Scanln(&nums[i])
	//}
	//QuickSort(nums, 0, len(nums)-1)
	//fmt.Println(nums)
	str := bufio.NewScanner(os.Stdin)
	str.Scan()
	fmt.Println(str.Text())
	nums := strings.Split(str.Text(), " ")

	fmt.Println(nums)
}
