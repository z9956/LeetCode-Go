package linked_list

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	cur := head

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	if head.Val == val {
		return head.Next
	}

	return head
}
