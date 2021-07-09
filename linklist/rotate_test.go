package linklist

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}

	if head == nil {
		return nil
	}

	cur := head
	for index := 0; index < k ; index++  {
		if cur.Next == nil {
			cur = head
		} else {
			cur = cur.Next
		}
	}

	if cur == nil {
		return head
	}

	kNode := cur

	for cur.Next != nil {
		cur = cur.Next
	}

	ret := kNode.Next
	if ret == nil {
		ret = kNode
	}

	cur.Next = head
	kNode.Next = nil

	return ret
}
