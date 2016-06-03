package atomic

import (
	"testing"
)

func checkBoolValue(value, target bool, t *testing.T) {
	if value != target {
		t.Fatalf("Value should be %t, but found %t instead", target, value)
	}
}

func TestAtomicBool(t *testing.T) {
	t.Log("Test: AtomicBool ...")

	ab := NewAtomicBool(false)
	checkBoolValue(ab.Get(), false, t)

	ab.Set(true)
	checkBoolValue(ab.Get(), true, t)

	t.Log("Passed ...")
}
