package array

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	dummyNode := ListNode{}
	endOfLastGroup := &dummyNode

	for head != nil {
		// 1. 找到group中的最后一个点
		end := findEndInGroup(head, k)
		if end == nil {
			break
		}

		// 2. 记录下一个group的起始点
		nextGroupHead := end.Next
		end.Next = nil
		// 3. 翻转链表
		headOfCurGroup := reverseList(head)
		// 4. 更改指向group的头尾节点
		endOfLastGroup.Next = headOfCurGroup
		head.Next = nextGroupHead

		endOfLastGroup = head
		head = nextGroupHead
	}

	return dummyNode.Next
}


func findEndInGroup(head *ListNode, k int)*ListNode  {
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		k--
		if k == 0 {
			return cur
		}

		cur = cur.Next
	}

	return nil
}
