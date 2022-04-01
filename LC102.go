package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		temp := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			ans[j] = append(ans[j], q[j].Val)
			if q[j].Left != nil {
				temp = append(temp, q[j].Left)
			}
			if q[j].Right != nil {
				temp = append(temp, q[j].Right)
			}
		}
		q = temp
	}
	return ans
}
