package linklist2

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 归并排序
	// 找到中点
	mid := findMid(head)

	// 对左右两边分别进行排序
	rightList := sortList(mid.Next)
	mid.Next = nil
	leftList := sortList(head)
	// 合并两个链表

	return mergeTwoList(leftList, rightList)
}

func mergeTwoList(first *ListNode, second *ListNode) *ListNode {
	if first == nil {
		return second
	}

	if second == nil {
		return first
	}

	dummyNode := ListNode{}
	curOfDummy := dummyNode.Next

	curFirst := first
	curSecond := second

	for curFirst != nil && curSecond != nil {
		if curFirst.Val < curSecond.Val {
			curOfDummy.Next = curFirst

			curOfDummy = curOfDummy.Next
			curFirst = curFirst.Next
		} else {
			curOfDummy.Next = curSecond

			curOfDummy = curOfDummy.Next
			curSecond.Next = curSecond.Next
		}

		curFirst = curFirst.Next
		curSecond = curSecond.Next
	}

	if curFirst != nil {
		curOfDummy.Next = curFirst
	} else {
		curOfDummy.Next = curSecond
	}

	return dummyNode.Next
}

func findMid(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	faster := head.Next

	for faster != nil && faster.Next != nil {
		faster = faster.Next.Next
		slow = slow.Next
	}

	return slow
}
