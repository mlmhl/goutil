package container

import "container/list"

type Hashable interface {
	HashCode() uint64

	Equal(other Hashable) bool
}

const (
	default_factor   = 0.75
	default_capacity = 16
)

type linkNode struct {
	Key Hashable
	Value interface{}
	Next *linkNode
}

type HashMap struct {
	size     uint32
	factor   float32
	capacity uint32

	buckets []*linkNode
}

//TODO adjust capacity to next power of 2.

func NewHashMap() *HashMap {
	return NewHashMapWithCapacityAndFactor(default_capacity, default_factor)
}

func NewHashMapWithCapacity(cap int32) *HashMap {
	return NewHashMapWithCapacityAndFactor(cap, default_factor)
}

func NewHashMapWithCapacityAndFactor(cap int32, factor float32) *HashMap {
	if cap < default_capacity {
		cap = default_capacity
	}
	return &HashMap{
		size:     0,
		factor:   factor,
		capacity: cap,

		buckets: make([]*linkNode, cap),
	}
}

func (hashMap *HashMap) Get(key Hashable) interface{} {
	node := hashMap.get(key)
	if node != nil {

	}
}

func (HashMap *HashMap) get(key Hashable) *linkNode {

}

func (hashMap *HashMap) Put(key Hashable, value interface{}) {
	if hashMap.size+1 > hashMap.capacity*hashMap.factor {
		hashMap.rehash()
	}
	hashMap.size++
	hashMap.put(key, &linkNode{
		Key: key,
		Value: value,
	})
}

func (hashMap *HashMap) put(key Hashable, node *linkNode) {
	index := key.HashCode() % hashMap.capacity
	if hashMap.buckets[index] == nil {
		hashMap.buckets[index] = node
	}
	node.Next, hashMap.buckets[index] = hashMap.buckets[index], node
}

func (hashMap *HashMap) rehash() {
	oldBuckets := hashMap.buckets

	hashMap.capacity = hashMap.capacity << 1
	hashMap.buckets = make([]*list.List, hashMap.capacity)

	for _, node := range oldBuckets {
		for node != nil {
			next := node.Next
			hashMap.put(node.Key, node)
			node = next
		}
	}
}
