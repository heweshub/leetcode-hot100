package tree

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func postorderTraversal(root *TreeNode) (ans []int) {
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		ans = append(ans, node.Val)
	}
	postorder(root)
	return
}
