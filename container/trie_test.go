package container

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	values := map[string]int{
		"asdfg": 1,
		"asdzxcv": 2,
		"qwerstyu": 3,
		"qweds": 4,
		"qwersmjkn": 5,
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

	t.Log("Test Trie: Iterator ...")
	iterator := trie.Iterator()
	for iterator.HasNext() {
		v := iterator.Next().(*Entry).GetValue()
		exist := false
		for _, t := range(values) {
			if v == t {
				exist = true
				break
			}
		}
		if !exist {
			t.Fatalf("Not exist value: %d", v)
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
