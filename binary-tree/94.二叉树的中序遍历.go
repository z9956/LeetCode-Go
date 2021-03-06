package binary_tree

import "github.com/z9956/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) (values []int) {
	if root == nil {
		return
	}

	values = append(values, inorderTraversal(root.Left)...)
	values = append(values, root.Val)
	values = append(values, inorderTraversal(root.Right)...)
	return
}
