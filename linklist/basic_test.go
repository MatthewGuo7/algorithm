package linklist

type Node struct {
	Next *Node
	Data int
}

type List struct {
	Header Node
}

func (r *List)Search(val int) *Node {
	cur := r.Header.Next

	for cur != nil {
		if cur.Data == val {
			return cur
		}
		cur = cur.Next
	}

	return nil
}

func (r *List) Insert(val int)  {
	newNode := &Node{
		Next: nil,
		Data: val,
	}

	prev := &r.Header
	cur := prev.Next

	for cur != nil {
		if cur.Data != val {
			prev = cur
			cur = cur.Next
			continue
		} else {
			break
		}
	}

	newNode.Next = prev.Next
	prev.Next = newNode
}