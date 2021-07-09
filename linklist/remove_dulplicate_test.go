package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}

	cur := dummyNode

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			val := cur.Next.Val
			for cur.Next != nil {
				if cur.Next.Val == val {
					cur.Next = cur.Next.Next
				} else {
					break
				}
			}
		} else {
			cur = cur.Next
		}
	}
	return dummyNode.Next
}
