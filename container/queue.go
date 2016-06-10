package container

import (
	"container/list"
	"fmt"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (queue *Queue) Len() int {
	return queue.list.Len()
}

func (queue *Queue) Push(v interface{}) {
	queue.list.PushBack(v)
}

func (queue *Queue) Front() interface{} {
	return queue.list.Front().Value
}

func (queue *Queue) Pop() {
	queue.list.Remove(queue.list.Front())
}

func (queue *Queue) String() string {
	buffer := []rune{'['}
	for element := queue.list.Front(); element != nil; element = element.Next() {
		buffer = append(buffer, ([]rune)(fmt.Sprintf("%v,", element.Value))...)
	}
	if buffer[len(buffer)-1] == ',' {
		buffer[len(buffer)-1] = ']'
	} else {
		buffer = append(buffer, ']')
	}
	return string(buffer)
}