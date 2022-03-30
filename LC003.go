package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	seen := make(map[byte]int)
	left := 0
	max := 0
	cur := 0
	for i := 0; i < n; i++ {
		cur++
		for {
			if _, ok := seen[s[i]]; ok {
				delete(seen, s[left])
				left++
				cur--
			} else {
				break
			}
		}
		if cur > max {
			max = cur
		}
		seen[s[i]]++
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
