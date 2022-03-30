package main

import "fmt"

func longestPalindrome1(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	max := 1
	begin := 0
	for Length := 2; Length <= n; Length++ {
		for i := 0; i < n; i++ {
			j := Length + i - 1
			if j >= n {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > max {
				max = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+max]
}

func longestPalindrome2(s string) string {
	n := len(s)
	start, end := 0, 0
	for i := 0; i < n; i++ {
		start1, end1 := find(s, i, i)
		start2, end2 := find(s, i, i+1)
		if end-start < end1-start1 {
			start, end = start1, end1
		}
		if end-start < end2-start2 {
			start, end = start2, end2
		}
	}
	return s[start : end+1]
}
func find(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		r++
		l--
	}
	return l + 1, r - 1
}

func main() {
	fmt.Println(longestPalindrome1("aa"))
}
