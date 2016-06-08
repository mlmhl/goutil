package container

import (
	"testing"
)

type testKey struct {
	id uint32
}

func newTestKey(id uint32) *testKey {
	return &testKey{
		id: id,
	}
}

func (key *testKey) HashCode() uint32 {
	return key.id
}

func (key *testKey) Equal(other Hashable) bool {
	if otherKey, ok := other.(*testKey); !ok {
		return false
	} else {
		return key.id == otherKey.id
	}
}

func checkHashMapContent(hashMap *HashMap, target string, t *testing.T) {
	content := hashMap.String()
	if target != content {
		t.Fatalf("Wrong content, Wanted %s, got %s", target, content)
	}
}

func TestHashMap(t *testing.T) {
	t.Log("Test: HashMap ...")

	hashMap := NewHashMap()

	t.Log("Test: Put ...")

	hashMap.Put(newTestKey(1), "1")
	hashMap.Put(newTestKey(2), "2")
	hashMap.Put(newTestKey(3), "3")

	t.Log(hashMap.String())
}