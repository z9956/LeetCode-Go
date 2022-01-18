package linked_list

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	less := &ListNode{0, nil}
	greater := &ListNode{0, nil}

	lessTail := less
	greaterTail := greater

	cur := head

	for cur != nil {
		if cur.Val < x {
			less.Next = cur
			less = less.Next
		} else {
			greater.Next = cur
			greater = greater.Next
		}

		cur = cur.Next
	}

	greater.Next = nil
	less.Next = greaterTail.Next
	return lessTail.Next
}
