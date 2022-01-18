package linked_list

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if k == 0 {
		return head
	}

	var length int = 1
	cur := head

	for cur.Next != nil {
		cur = cur.Next
		length++
	}

	var move = length - (k % length)
	cur.Next = head

	for move > 0 {
		cur = cur.Next
		move--
	}

	ans := cur.Next

	cur.Next = nil
	return ans
}
