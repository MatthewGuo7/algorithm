package linklist2

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if k <= 0 {
		return head
	}

	length := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		length++
	}

	cur.Next = head
	step := length - k%length - 1

	cur = head
	for index := 0; index <= step; index++ {
		cur = cur.Next
	}

	ret := cur.Next
	cur.Next = nil

	return ret
}
