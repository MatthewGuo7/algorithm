package linklist2

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	head = copyNewNode(head)
	copyRandomPointer(head)
	return splitList(head)
}

func splitList(head *Node) *Node {
	if head == nil {
		return head
	}

	dummyNode := Node{}

	cur := head
	next := head.Next

	curOfDummy := &dummyNode

	for next != nil {
		cur.Next = next.Next

		cur = next.Next

		curOfDummy.Next = next
		curOfDummy = curOfDummy.Next

		if cur == nil {
			break
		}

		next = cur.Next
	}

	curOfDummy.Next = nil

	return dummyNode.Next
}

func copyRandomPointer(head *Node) {
	if head == nil {
		return
	}

	cur := head

	for cur != nil {
		if cur.Next.Random != nil {
			cur.Next.Random = cur.Random.Next
		}

		cur = cur.Next.Next
	}
}

func copyNewNode(head *Node) *Node {
	dummyNode := Node{Next: head}

	cur := dummyNode.Next

	for cur != nil {
		newNode := *cur

		newNode.Next = cur.Next
		cur.Next = &newNode

		cur = cur.Next.Next
	}

	return dummyNode.Next
}
