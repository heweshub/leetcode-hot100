package main

import "fmt"

func hammingWeight(num uint32) (ans int) {
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			fmt.Println(1<<i, 1<<i&num)
			//fmt.Println(1 << i & num)
			ans++
		}
	}
	return
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
