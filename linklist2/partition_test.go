package linklist2

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	dummyLeft := ListNode{}
	left := &dummyLeft

	dummyRight := ListNode{}
	right := &dummyRight

	cur := head

	for cur != nil {
		if cur.Val < x {
			left.Next = cur
			left = left.Next
		} else {
			right.Next = cur
			right = right.Next
		}

		cur = cur.Next
	}

	right.Next = nil
	left.Next = dummyRight.Next

	return left.Next
}
