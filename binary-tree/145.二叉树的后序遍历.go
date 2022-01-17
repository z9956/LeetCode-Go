package binary_tree

func postorderTraversal(root *TreeNode) (values []int) {
	if root == nil {
		return
	}

	values = append(values, postorderTraversal(root.Left)...)
	values = append(values, postorderTraversal(root.Right)...)
	values = append(values, root.Val)
	return
}
