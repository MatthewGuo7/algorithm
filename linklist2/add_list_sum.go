package linklist2

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := ListNode{}
	curOfDummy := &dummyNode

	curOfL1 := l1
	curOfL2 := l2

	carry := 0
	for curOfL1 != nil || curOfL2 != nil {
		value := carry
		if curOfL1 != nil {
			value += curOfL1.Val
			curOfL1 = curOfL1.Next
		}

		if curOfL2 != nil {
			value += curOfL2.Val
			curOfL2 = curOfL2.Next
		}

		carry = value / 10
		value = value % 10

		newNode := &ListNode{Val: value}
		curOfDummy.Next = newNode

		curOfDummy = curOfDummy.Next
	}

	if carry > 0 {
		curOfDummy.Next = &ListNode{
			Val: carry,
		}
	}

	return dummyNode.Next
}
