package container

import (
	"fmt"
)

type Hashable interface {
	HashCode() uint32

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

func NewHashMapWithCapacity(cap uint32) *HashMap {
	return NewHashMapWithCapacityAndFactor(cap, default_factor)
}

func NewHashMapWithCapacityAndFactor(cap uint32, factor float32) *HashMap {
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

func (hashMap *HashMap) Size() uint32 {
	return hashMap.size
}

func (hashMap *HashMap) Get(key Hashable) interface{} {
	node := hashMap.get(key)
	if node != nil {
		return node.Value
	} else {
		return nil
	}
}

func (hashMap *HashMap) Put(key Hashable, value interface{}) {


	if (float32)(hashMap.size+1) > (float32)(hashMap.capacity)*hashMap.factor {
		hashMap.rehash()
	}
	hashMap.size++
	hashMap.put(key, &linkNode{
		Key: key,
		Value: value,
	})
}

func (hashMap *HashMap) Remove(key Hashable) {
	index := key.HashCode() % hashMap.capacity
	node := hashMap.buckets[index]

	if node == nil {
		return
	} else if node.Key.Equal(key) {
		hashMap.buckets[index] = node.Next
	} else {
		for node.Next != nil {
			if node.Next.Key.Equal(key) {
				node.Next = node.Next.Next
			}
			node = node.Next
		}
	}
}

func (hashMap *HashMap) String() string {
	buffer := []rune{'['}

	for _, node := range hashMap.buckets {
		for node != nil {
			buffer = append(buffer,
				([]rune)(fmt.Sprintf("{%v: %v},", node.Key, node.Value))...)
		}
	}

	if buffer[len(buffer)-1] == ',' {
		buffer[len(buffer)-1] = ']'
	} else {
		buffer = append(buffer, ']')
	}

	return (string)(buffer)
}

func (hashMap *HashMap) get(key Hashable) *linkNode {
	node := hashMap.buckets[key.HashCode() % hashMap.capacity]
	for node != nil {
		if node.Key.Equal(key) {
			return node
		}
		node = node.Next
	}
	return nil
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
	hashMap.buckets = make([]*linkNode, hashMap.capacity)

	for _, node := range oldBuckets {
		for node != nil {
			next := node.Next
			hashMap.put(node.Key, node)
			node = next
		}
	}
}
