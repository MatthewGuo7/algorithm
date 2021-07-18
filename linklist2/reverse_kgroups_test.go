package linklist2

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	dummyNode := ListNode{Next: head}

	for prev := &dummyNode; prev != nil; {
		prev = reverseKNodes(prev, k)
	}

	return dummyNode.Next
}

func reverseKNodes(prev *ListNode, k int) *ListNode {
	// 找到第k个点

	cur := prev

	for index := 0; index < k; index++ {
		cur = cur.Next
		if cur == nil {
			return nil
		}
	}


	first := prev.Next

	dummyNode := ListNode{Next: cur.Next}

	cur.Next = nil
	headOfList := reverseList(prev.Next)

	prev.Next = headOfList
	first.Next = dummyNode.Next

	return first
}
