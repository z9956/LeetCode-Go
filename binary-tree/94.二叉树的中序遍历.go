package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (values []int) {
	if root == nil {
		return
	}

	values = append(values, inorderTraversal(root.Left)...)
	values = append(values, root.Val)
	values = append(values, inorderTraversal(root.Right)...)
	return
}
