package stack_queue

import "container/list"

func sortStack(inputStack *list.List) *list.List {
	outputStack := list.New()

	for inputStack.Len() > 0 {
		top := inputStack.Back()
		inputStack.Remove(top)

		v := top.Value.(int)
		for outputStack.Len() > 0 && outputStack.Back().Value.(int) < v {
			topInOutputStack := outputStack.Back()
			outputStack.Remove(topInOutputStack)
			inputStack.PushBack(topInOutputStack)
		}

		outputStack.PushBack(top)
	}

	return outputStack
}
