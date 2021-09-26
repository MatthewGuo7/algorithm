package linklist2

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}

	mid := findMid(head)
	right := reverseList(mid.Next)
	mid.Next = nil

	ret := mergeList(head, right)

	head = ret
}

func mergeList(first *ListNode, second *ListNode) *ListNode {
	dummyNode := ListNode{}
	cur := &dummyNode

	for first != nil && second != nil {
		cur.Next = first
		cur = cur.Next

		cur.Next = second
		cur = cur.Next

		first = first.Next
		second = second.Next
	}

	if first != nil {
		cur.Next = first
	} else  {
		cur.Next = second
	}

	return dummyNode.Next
}


