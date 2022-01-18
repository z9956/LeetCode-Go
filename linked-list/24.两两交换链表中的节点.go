package linked_list

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var next = head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
