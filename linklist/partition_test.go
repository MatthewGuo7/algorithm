package linklist

func partition(head *ListNode, x int) *ListNode {
	dummyLeft := &ListNode{
		Val:  0,
		Next: nil,
	}

	dummyRight := &ListNode{
		Val:  0,
		Next: nil,
	}

	left := dummyLeft
	right := dummyRight

	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}

	right.Next = nil
	left.Next = dummyRight.Next

	return dummyLeft.Next
}
