package linked_list

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	cur := list1

	for i := 0; i < a-1; i++ {
		cur = cur.Next
	}

	node := cur.Next
	for i := a; i < b+1; i++ {
		node = node.Next
	}

	cur.Next = list2
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = node

	return list1
}
