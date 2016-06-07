package container

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Log("Test: Stack ...")

	stack := NewStack()
	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		stack.Push(value)
	}
	for i := len(values) - 1; i >= 0; i-- {
		v := stack.Peek()
		stack.Pop()
		if v != values[i] {
			t.Fatalf("Wrong value, Wanted %d, got %d", values[i], v)
		}
	}

	t.Log("Passed")
}
