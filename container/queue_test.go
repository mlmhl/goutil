package container

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Log("Test: Queue ...")

	queue := NewQueue()
	values := []int{1, 2, 3, 4, 5}

	for _, v := range values {
		queue.Push(v)
	}
	pos := 0
	for queue.Len() > 0 {
		v := queue.Front()
		if v.(int) != values[pos] {
			t.Fatalf("Wrong value(%d), Wanted %d, got %d", pos, values[pos], v)
		}
		pos++
		queue.Pop()
	}

	t.Log("Passed")
}
