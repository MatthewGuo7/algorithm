package linklist2

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummyNode := &ListNode{Next: head}
	cur := dummyNode

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			valueToDelete := cur.Next.Val
			for cur.Next != nil {
				if cur.Next.Val == valueToDelete {
					cur.Next = cur.Next.Next
				 } else {
				 	break
				}
			}
		} else {
			cur = cur.Next
		}

	}

	return dummyNode.Next
}
