package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	//当前层的队列
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		// 定义下一层的队列
		p := []*TreeNode{}
		// 循环遍历当前层的所有元素，并添加子树到下一层队列p中
		for j := 0; j < len(q); j++ {
			node := q[i]
			// 结果集中添加当前层的结点值
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}
