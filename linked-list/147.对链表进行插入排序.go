package linked_list

func insertionSortList(head *ListNode) *ListNode {
	cur := head
	var sorted = &ListNode{0, nil}

	for cur != nil {
		next := cur.Next
		sortedCur := sorted

		for sortedCur.Next != nil && sortedCur.Next.Val < cur.Val {
			sortedCur = sortedCur.Next
		}

		cur.Next = sortedCur.Next
		sortedCur.Next = cur

		cur = next
	}

	return sorted.Next
}
