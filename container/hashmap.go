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
	Key   Hashable
	Value interface{}
	Next  *linkNode
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

func (hashMap *HashMap) Len() uint32 {
	return hashMap.size
}

func (hashMap *HashMap) Contains(key Hashable) bool {
	return hashMap.Get(key) != nil
}

func (hashMap *HashMap) Get(key Hashable) interface{} {
	node := hashMap.buckets[key.HashCode()%hashMap.capacity]
	for node != nil {
		if node.Key.Equal(key) {
			return node.Value
		}
		node = node.Next
	}
	return nil
}

func (hashMap *HashMap) Put(key Hashable, value interface{}) {
	code := key.HashCode()
	index := code % hashMap.capacity

	// If key exist, update its value.
	for node := hashMap.buckets[index]; node != nil; {
		if node.Key.Equal(key) {
			node.Value = value
			return
		}
		node = node.Next
	}

	hashMap.size++
	node := &linkNode{
		Key:   key,
		Value: value,
	}
	if (float32)(hashMap.size) > (float32)(hashMap.capacity)*hashMap.factor {
		hashMap.rehash()
		hashMap.put(code, node)
	} else {
		hashMap.insert((int)(index), node)
	}
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
			node = node.Next
		}
	}

	fmt.Println("mark")

	if buffer[len(buffer)-1] == ',' {
		buffer[len(buffer)-1] = ']'
	} else {
		buffer = append(buffer, ']')
	}

	return (string)(buffer)
}

func (hashMap *HashMap) put(code uint32, node *linkNode) {
	hashMap.insert((int)(code%hashMap.capacity), node)
}

func (hashMap *HashMap) insert(index int, node *linkNode) {
	node.Next, hashMap.buckets[index] = hashMap.buckets[index], node
}

func (hashMap *HashMap) rehash() {
	oldBuckets := hashMap.buckets

	hashMap.capacity = hashMap.capacity << 1
	hashMap.buckets = make([]*linkNode, hashMap.capacity)

	for _, node := range oldBuckets {
		for node != nil {
			next := node.Next
			hashMap.put(node.Key.HashCode(), node)
			node = next
		}
	}
}
