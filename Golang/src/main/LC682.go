package main

import (
	"strconv"
)

func calPoints(ops []string) int {
	var score []int
	for _, v := range ops {
		if num, err := strconv.Atoi(v); err == nil {
			score = append(score, num)
		} else if v == "C" {
			score = score[:len(score)-1]
		} else if v == "D" {
			score = append(score, score[len(score)-1]*2)
		} else if v == "+" {
			score = append(score, score[len(score)-1]+score[len(score)-2])
		}
	}
	sum := 0
	for i := 0; i < len(score); i++ {
		sum += score[i]
	}
	return sum
}

func main() {
	calPoints([]string{"1", "2"})
}
