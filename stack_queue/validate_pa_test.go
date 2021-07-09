package stack_queue

import (
	"container/list"
	"fmt"
	"testing"
)

func isLeftP(v byte) bool {
	return v == '(' || v == '{' || v == '['
}

func matchedLeft(left, right byte) bool {
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

func isValid(s string) bool {
	stack := list.New()

	for _, v := range s {
		if isLeftP(byte(v)) {
			stack.PushBack(v)
		} else {
			//是右括号
			if stack.Len() == 0 {
				return false
			}

			top := stack.Back()
			if !matchedLeft(byte(top.Value.(int32)), byte(v)) {
				return false
			}

			stack.Remove(top)

		}
	}

	return stack.Len() == 0
}

func TestValid(t *testing.T) {
	s := "()"
	fmt.Println(isValid(s))
}
