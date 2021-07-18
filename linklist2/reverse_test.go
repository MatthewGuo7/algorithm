package linklist2

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyNode := ListNode{}

	cur := head
	next := cur.Next

	for cur != nil {
		next = cur.Next

		cur.Next = dummyNode.Next
		dummyNode.Next = cur

		cur = next
	}

	return dummyNode.Next
}

func reverseList2(head *ListNode) *ListNode  {
	dummyNode := ListNode{}

	reverseListHelper2(head, &dummyNode)

	return dummyNode.Next
}

func reverseListHelper2(head *ListNode, dummyNode *ListNode)  {
	if head == nil {
		return
	}

	next := head.Next

	head.Next = dummyNode.Next
	dummyNode.Next = head

	reverseListHelper2(next, dummyNode)
}
