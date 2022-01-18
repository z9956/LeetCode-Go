package linked_list

import (
	"github.com/z9956/LeetCode-Go/structures"
)

type ListNode = structures.ListNode

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
