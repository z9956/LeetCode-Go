package linked_list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	i := headA
	j := headB

	for i != j {
		if i == nil {
			i = headA
		} else {
			i = i.Next
		}

		if j == nil {
			j = headB
		} else {
			j = j.Next
		}

	}

	return i
}
