package tree

type item struct {
	idx int
	*TreeNode
}

//借助了层次遍历的思想，加了一个索引的概念
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans, q := 1, []item{{0, root}}
	for len(q) > 0 {
		if l := q[len(q)-1].idx - q[0].idx + 1; l > ans {
			ans = l
		}

		temp := []item{}
		for _, p := range q {
			if p.Left != nil {
				temp = append(temp, item{p.idx * 2, p.Left})
			}
			if p.Right != nil {
				temp = append(temp, item{p.idx*2 + 1, p.Right})
			}
		}
		q = temp
	}
	return ans
}
