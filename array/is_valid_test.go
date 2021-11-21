package array

import "container/list"

func isValid(s string) bool {
	stack := list.New()

	for _, b := range s {
		if isLeftP(byte(b)) {
			stack.PushBack(b)
		} else {
			if stack.Len() == 0 {
				return false
			}

			node := stack.Back()
			left := node.Value.(int32)

			if matchLeft(byte(left), byte(b)) {
				stack.Remove(node)
			} else {
				return false
			}
		}
	}

	return stack.Len() == 0
}

func isLeftP(b byte) bool  {
	return b == '(' || b == '[' || b == '{'
}

func matchLeft(left, right byte) bool  {
	switch right {
	case ')':
		return left == '(' 
	case '}':
		return left == '{'
	case ']':
		return left == '['
	}

	return false
}
