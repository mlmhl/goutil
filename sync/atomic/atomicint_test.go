package atomic

import (
	"testing"
)

func checkIntValue(value, target int64, t *testing.T) {
	if value != target {
		t.Fatalf("Wanted %d, got %d", target, value)
	}
}

func TestOperation(t *testing.T) {
	t.Log("Test: AtomicInt ...")

	var value int64 = 10
	n := NewAtomicInt(value)
	checkIntValue(n.Get(), value, t)

	n.Add(10)
	value += 10
	checkIntValue(n.Get(), value, t)

	value = 100
	n.Set(value)
	checkIntValue(n.Get(), value, t)

	value -= 10
	n.Add(-10)
	checkIntValue(n.Get(), value, t)

	n.Reset()
	checkIntValue(n.Get(), 0, t)

	t.Log("Passed ...")
}
