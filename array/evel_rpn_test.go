package array

import (
	"container/list"
	"errors"
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	stack := list.New()

	for _, token := range tokens {
		ok := isEval(token)
		if !ok {
			value, err := strconv.Atoi(token)
			if err != nil {
				return 0
			}

			stack.PushBack(value)
			continue
		}


		firstNode := stack.Back()
		first := firstNode.Value.(int)
		stack.Remove(firstNode)

		secondNode := stack.Back()
		second := secondNode.Value.(int)
		stack.Remove(secondNode)

		result, err := calc(second, first, token)
		if err != nil {
			return 0
		}

		stack.PushBack(result)

	}

	if stack.Len() == 1 {
		ret := stack.Back().Value.(int)
		return ret
	}

	return 0
}

func calc(first, second int, token string) (int, error) {
	switch token {
	case "+":
		return first + second, nil
	case "-":
		return first - second, nil
	case "*":
		return first * second, nil
	case "/":
		return first / second, nil
	}

	return 0, errors.New("invalid token")
}

func isEval(token string) bool {
	return (token == "+" || token == "-" || token == "*" || token == "/")
}
