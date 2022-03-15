package main

import "fmt"

// 3. 无重复字符的最长子串

func lengthOfLongestSubstrings(s string) int {
	// 定义map映射
	m := map[byte]int{}
	n := len(s)
	// 右指针
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 映射中删除前一个起点
			delete(m, s[i-1])
		}
		// 判断条件是未出现过并且右指针不出界
		// 直到找到已经出现的元素或者到字符串结尾
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		// 比较当前最长无重复字符串与之前的比较
		ans = max(ans, rk-i+1)
	}
	fmt.Println(m)
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(lengthOfLongestSubstrings("pwwkew"))
}
