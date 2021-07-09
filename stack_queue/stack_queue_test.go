package stack_queue

import "container/list"

type StackWithMax struct {
	ValueStack *list.List
	MaxStack   *list.List
}

func NewStackWithMax() *StackWithMax {
	return &StackWithMax{
		ValueStack: list.New(),
		MaxStack:   list.New(),
	}
}

func (r *StackWithMax) GetMax() int {
	return r.MaxStack.Back().Value.(int)
}

func (r *StackWithMax) Push(value int) {
	if r.MaxStack.Len() == 0 || r.MaxStack.Back().Value.(int) <= value {
		r.MaxStack.PushBack(value)
	}

	r.ValueStack.PushBack(value)
}

func (r *StackWithMax) Pop() int {
	top := r.ValueStack.Back()
	r.ValueStack.Remove(top)

	topInMax := r.MaxStack.Back()
	if top.Value.(int) == topInMax.Value.(int) {
		r.MaxStack.Remove(topInMax)
	}

	r.ValueStack.Remove(top)
	return top.Value.(int)
}
