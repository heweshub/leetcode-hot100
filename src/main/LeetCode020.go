package main

import "fmt"

func isVaild(s string) bool {
	var stack []byte
	mapDict := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == mapDict[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	s := "}[[]]"
	fmt.Println(isVaild(s))
}
