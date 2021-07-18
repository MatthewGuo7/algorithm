package linklist2

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur := head

	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			continue
		}

		cur = cur.Next
	}

	return head
}
