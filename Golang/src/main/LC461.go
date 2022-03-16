package main

import "fmt"

// 得到x和y的异或结果之后，逐位与1相&与，结果相加
func hammingDistance(x, y int) (ans int) {
	for i := x ^ y; i > 0; i >>= 1 {
		ans += i & 1
	}
	return
}

func main() {
	x := 1
	y := 4
	fmt.Println(x ^ y)
	fmt.Println(hammingDistance(x, y))
}
