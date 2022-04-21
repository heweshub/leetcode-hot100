package main

import "fmt"

func countOdds(low int, high int) int {
	cnt := high - low
	if cnt&1 == 1 {
		return (cnt + 1) >> 1
	} else {
		if low&1 == 1 {
			return cnt>>1 + 1
		} else {
			return cnt >> 1
		}
	}
}

func main() {
	fmt.Println(countOdds(3, 7))
	fmt.Println(countOdds(8, 10))
}
