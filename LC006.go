package main

import "fmt"

func convert(s string, r int) string {
	n := len(s)
	if r == 1 || r >= n {
		return s
	}
	// 可以发现周期t
	t := 2*r - 2
	ans := make([]byte, 0, n)
	// 循环r次
	for i := 0; i < r; i++ {
		for j := 0; j+i < n; j += t {
			// 当前周期中的第一个数
			ans = append(ans, s[j+i])
			if i > 0 && i < r-1 && j+t-i < n {
				// 当前周期中的第二个数
				ans = append(ans, s[j+t-i])
			}
		}
	}
	return string(ans)
}

func main() {
	s := "PAYPALISHIRING"
	fmt.Println(convert(s, 3))
	fmt.Println(convert(s, 4))
}
