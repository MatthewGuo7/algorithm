package array

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	dummyNode := ListNode{}

	for cur != nil {
		next := cur.Next

		cur.Next = dummyNode.Next
		dummyNode.Next = cur

		cur = next
	}

	return dummyNode.Next
}
