package main

import "fmt"

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(lmain, rmain int, path string)
	dfs = func(lmain, rmain int, path string) {
		if len(path) == 2*n {
			res = append(res, path)
			return
		}
		if lmain > 0 {
			dfs(lmain-1, rmain, path+"(")
		}
		if lmain < rmain {
			dfs(lmain, rmain-1, path+")")
		}
	}
	dfs(n, n, "")
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
