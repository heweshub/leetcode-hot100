package main

import "fmt"

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "xzy",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	// 递归结束的条件
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		// 取到每一层的数字
		digit := string(digits[index])
		// 数字在电话按键上对应的字母
		letters := phoneMap[digit]
		lettersCount := len(letters)
		// 遍历该数字对应的所有字母
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
