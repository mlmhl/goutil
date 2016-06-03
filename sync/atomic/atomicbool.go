package atomic

import (
	"sync"
)

type AtomicBool struct {
	value bool
	lock sync.Mutex
}

// Get a new AtomicBool with value "value".
func NewAtomicBool(value bool) *AtomicBool {
	return &AtomicBool{
		value: value,
		lock: sync.Mutex{},
	}
}

// Set value.
func (ab *AtomicBool) Set(value bool) {
	ab.lock.Lock()
	defer ab.lock.Unlock()
	ab.value = value
}

// Get value.
func (ab *AtomicBool) Get() bool {
	ab.lock.Lock()
	defer ab.lock.Unlock()
	return ab.value
}