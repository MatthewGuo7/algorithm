package stack_queue

import "container/list"

type QueueByStack struct {
	InputStack  *list.List
	OutputStack *list.List
}

func NewQueueByStack() *QueueByStack {
	return &QueueByStack{
		InputStack:  list.New(),
		OutputStack: list.New(),
	}
}

func (r *QueueByStack)enqueue(value int)  {
	r.InputStack.PushBack(value)
}

func (r *QueueByStack) dequeue() int {
	if r.OutputStack.Len() > 0 {
		top := r.OutputStack.Back()
		r.OutputStack.Remove(top)
		return top.Value.(int)
	}

	for r.InputStack.Len() > 0 {
		top := r.InputStack.Back()
		r.OutputStack.PushBack(top)
		r.InputStack.Remove(top)
	}

	ret := r.OutputStack.Back()
	return ret.Value.(int)
}
