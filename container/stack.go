package container

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{
		list: list.New(),
	}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Peek() interface{} {
	return stack.list.Back().Value
}

func (stack *Stack) Pop() {
	stack.list.Remove(stack.list.Back())
}

func (stack *Stack) IsEmpty() bool {
	return stack.list.Len() == 0
}
