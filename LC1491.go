package main

import (
	"fmt"
	sort2 "sort"
)

func average(salary []int) float64 {
	sort2.Ints(salary)
	return sum(salary[1 : len(salary)-1])
}

func sum(a []int) float64 {
	count := 0
	for i := range a {
		count += a[i]
	}
	return float64(count) / float64(len(a))
}

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
}
