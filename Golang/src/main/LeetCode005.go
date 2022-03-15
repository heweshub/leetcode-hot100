package main

import "fmt"

// 5. 最长回文子串的Golang版本实现
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for j := 1; j < n+1; j++ {
		for i := 0; i < n; i++ {
			k := j + i - 1
			if k >= n {
				break
			}
			if s[i] != s[k] {
				dp[i][k] = false
			} else {
				if k-i < 3 {
					dp[i][k] = true
				} else {
					dp[i][k] = dp[i+1][k-1]
				}
			}
			if dp[i][k] && k-i+1 > maxLen {
				maxLen = k - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func main() {
	fmt.Println(longestPalindrome("abasadkkdjjajdaww"))
}
