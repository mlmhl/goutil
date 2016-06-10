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

func checkHashMapValues(key *testKey, target string, value string, t *testing.T) {
	if target != value {
		t.Fatalf("Wrong value(%v): Wanted %s, got %s", key, target, value)
	}
}

func TestHashMap(t *testing.T) {
	hashMap := NewHashMap()

	keys := []*testKey{newTestKey(1), newTestKey(2),newTestKey(3)}
	values := []string{"1", "2", "3"}

	t.Log("Test: Put ...")
	for i := 0; i < len(keys); i++ {
		hashMap.Put(keys[i],values[i])
	}
	checkHashMapContent(hashMap, "[{&{1}: 1},{&{2}: 2},{&{3}: 3}]", t)
	t.Log("Passed")

	t.Log("Test: Get ...")
	for i := 0; i < len(keys); i++ {
		checkHashMapValues(keys[i], values[i], hashMap.Get(keys[i]).(string), t)
	}
	t.Log("Passed")

	t.Log("Test: Rehash ...")
	hashMap = NewHashMapWithCapacity(2)
	for i := 0; i < len(keys); i++ {
		hashMap.Put(keys[i],values[i])
	}
	checkHashMapContent(hashMap, "[{&{1}: 1},{&{2}: 2},{&{3}: 3}]", t)
	t.Log("Passed")

	t.Log("Test: Remove ...")
	hashMap.Remove(keys[0])
	hashMap.Remove(keys[2])
	checkHashMapContent(hashMap, "[{&{2}: 2}]", t)
	t.Log("Passed")

}