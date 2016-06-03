package atomic

import (
	"sync/atomic"
)

// Atomic integer.

type AtomicInt int64

func NewAtomicInt(n int64) *AtomicInt {

	num := new(int64)
	*num = n
	return (*AtomicInt)(num)
}

// Operations.

func (n *AtomicInt) Get() int64 {

	return atomic.LoadInt64((*int64)(n))
}

func (n *AtomicInt) Add(delta int64) {

	atomic.AddInt64((*int64)(n), delta)
}

func (n *AtomicInt) Set(value int64) {

	atomic.StoreInt64((*int64)(n), value)
}

func (n *AtomicInt) Reset() {

	atomic.StoreInt64((*int64)(n), 0)
}