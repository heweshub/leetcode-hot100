package dp

func minimumTotal(t [][]int) int {
	if len(t) <= 1 {
		return t[0][0]
	}
	cpt := t
	for i := 1; i < len(cpt); i++ {
		for j := 0; j < len(cpt[i]); j++ {
			if j > 0 && j < len(cpt[i])-1 {
				cpt[i][j] += min(cpt[i-1][j], cpt[i-1][j-1])
			} else if j > 0 {
				cpt[i][j] += cpt[i-1][j-1]
			} else if j == 0 {
				cpt[i][j] += cpt[i-1][0]
			}
		}
	}
	var minLen int
	for _, v := range cpt[len(cpt)-1] {
		if v < minLen {
			minLen = v
		}
	}
	return minLen
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
