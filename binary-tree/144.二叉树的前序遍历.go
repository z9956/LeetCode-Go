package binary_tree

func preorderTraversal(root *TreeNode) (values []int) {
	if root == nil {
		return
	}
	values = append(values, root.Val)
	values = append(values, preorderTraversal(root.Left)...)
	values = append(values, preorderTraversal(root.Right)...)
	return
}
