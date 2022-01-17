package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0

	cur := head

	for cur.Next != nil {
		cur = cur.Next
		length++
	}

	dummy := &ListNode{0, head}
	cur = dummy

	for i := 0; i <= length-n; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next

	return dummy.Next
}
