package linklist2

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k < 0 {
		return nil
	}

	faster := head
	for index := 0; index < k;index++{
		faster = faster.Next
	}

	if faster == nil {
		return nil
	}

	slower := head
	for faster != nil {
		faster = faster.Next
		slower = slower.Next
	}

	return slower
}
