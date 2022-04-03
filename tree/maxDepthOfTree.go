package tree

func maxDepth(root *TreeNode) (ans int) {
	var depthOfTree func(node *TreeNode, depth int)
	depthOfTree = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		if depth > ans {
			ans = depth
		}
		depthOfTree(node.Left, depth)
		depthOfTree(node.Right, depth)
	}
	depthOfTree(root, 0)
	return
}
