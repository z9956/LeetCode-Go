package linked_list

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil {
		if fast.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
