package encoding

import (
	"testing"
)

func TestInt64(t *testing.T) {
	numbers := []int64{1, 0, -1, (2 << 32) -1, -(2 << 32)}

	t.Log("Test: Int64...")

	for _, target := range(numbers) {
		v := DefaultDecoder.Int64(DefaultEncoder.Int64(target))
		if v != target {
			t.Fatalf("Wrong value: Wanted %d, got %d", target, v)
		}
	}

	t.Log("Passed...")
}