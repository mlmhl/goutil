package container

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	values := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}

	t.Log("Test Trie: Put ...")
	for k, v := range(values) {
		trie.Put(k, v)
	}
	if (trie.Size() != len(values)) {
		t.Fatalf("Wrong size: Wanted %d, got %d", len(values), trie.Size())
	}
	t.Log("Passed")

	t.Log("Test Trie: Get ...")
	for k, target := range(values) {
		v := trie.Get(k).(int)
		if v != target {
			t.Fatalf("Wrong value(%s): Wanted %d, got %d", k, target, v);
		}
	}
	t.Log("Passed")

	t.Log("Test Trie: Remove ...")
	i := 4
	for k, _ := range(values) {
		trie.Remove(k)
		if trie.Size() != i {
			t.Fatalf("Wrong size: Wanted %d, got %d", i, trie.Size())
		}
		i--
	}
	t.Log("Passed")
}
