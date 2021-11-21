package dandiao_stack_test

import "container/list"

func finalPrices(prices []int) []int {
	if len(prices) == 0 {
		return nil
	}

	retPrices := make([]int, len(prices))
	copy(retPrices, prices)

	stack := list.New()

	for i := 0; i < len(prices); i++ {
		curPrice := prices[i]
		for stack.Len() > 0 {
			node := stack.Back()
			top := node.Value.(int)

			if retPrices[top] >= curPrice {
				retPrices[top] = retPrices[top] - curPrice
				stack.Remove(node)
			}
		}

		stack.PushBack(i)
	}

	return retPrices
}
