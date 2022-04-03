package tree

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func prevorderTraversal(root *TreeNode) (ans []int) {
	var prevorder func(node *TreeNode)
	prevorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		prevorder(node.Left)
		prevorder(node.Right)
	}
	prevorder(root)
	return
}
