package linked_list

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	ans := &ListNode{0, head}
	cur := ans

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			val := cur.Next.Val

			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return ans.Next
}
