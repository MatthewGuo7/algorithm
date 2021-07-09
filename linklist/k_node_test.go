package linklist

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	for index := 0; index < k; index++{
		fast = fast.Next
	}

	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
