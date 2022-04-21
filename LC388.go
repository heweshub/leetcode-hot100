package main

func lengthLongestPath(input string) (ans int) {
	st := []int{}
	for i, n := 0, len(input); i < n; {
		depth := 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++
		}
		length, isFile := 0, false
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}
		i++

		for len(st) >= depth {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			length += st[len(st)-1] + 1
		}
		if isFile {
			ans = max(ans, length)
		} else {
			st = append(st, length)
		}
	}
	return
}
