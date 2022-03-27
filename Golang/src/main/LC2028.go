package main

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	summ := 0
	for i := 0; i < len(rolls); i++ {
		summ += rolls[i]
	}
	sumn := (n+len(rolls))*mean - summ
	if sumn < n || sumn > n*6 {
		return []int{}
	}
	avg := sumn / n
	res := []int{}
	for j := 0; j < n && sumn > avg; j++ {
		res = append(res, avg)
		sumn -= avg
	}
	res = append(res, sumn)
	return res
}

func main() {
	fmt.Println(missingRolls([]int{3, 2, 4, 3}, 4, 2))
}
