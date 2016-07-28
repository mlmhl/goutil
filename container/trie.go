package container

type node struct {
	value interface{}
	table map[rune]*node
}

func newNode(value interface{}) *node {
	return &node{
		value: value,
		table: map[rune]*node{},
	}
}

func (node *node) hasValue() bool {
	return node.value != nil
}

type Trie struct {
	size int
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode(nil),
	}
}

func (trie *Trie) Size() int {
	return trie.size
}

func (trie *Trie) Empty() bool {
	return trie.Size() == 0
}

// Get key's value, if not found return nil instead.
func (trie *Trie) Get(key string) interface{} {
	node := trie.root
	for _, r := range(key) {
		next, ok := node.table[r]
		if !ok {
			return nil
		}
		node = next
	}
	return node.value
}

func (trie *Trie) Contains(key string) bool {
	return trie.Get(key) != nil
}

// Put a key-value, if key has exist, its value will be override.
func (trie *Trie) Put(key string, value interface{}) {
	node := trie.root

	for _, r := range(key) {
		next, ok := node.table[r]
		if !ok {
			next = newNode(nil)
			node.table[r] = next
		}
		node = next
	}

	if node.value == nil {
		trie.size++
	}
	node.value = value
}

// Remove key's value is exist.
// Won't modify Trie's structure.
func (trie *Trie) Remove(key string) {
	node := trie.root

	for _, r := range(key) {
		next, ok := node.table[r]
		if !ok {
			break
		}
		node = next
	}

	if node.value != nil {
		node.value = nil
		trie.size--
	}
}

func (trie *Trie) Iterator() Iterator {
	return newTrieIterator(trie.root)
}

type trieIterator struct {
	stack *Stack
}

func newTrieIterator(node *node) *trieIterator {
	stack := NewStack()
	if node != nil {
		stack.Push(node)
	}
	return &trieIterator{
		stack: stack,
	}
}

func (iterator *trieIterator) HasNext() bool {
	iterator.rollToNextValue()
	return !iterator.stack.IsEmpty()
}

func (iterator *trieIterator) Next() interface{} {
	if !iterator.HasNext() {
		return nil
	}
	node := iterator.stack.Peek().(*node)
	iterator.stack.Pop()
	return node.value
}

func (iterator *trieIterator) rollToNextValue() {
	for {
		if iterator.stack.IsEmpty() {
			break
		}
		node := iterator.stack.Peek().(*node)
		if node.hasValue() {
			break
		}
		iterator.stack.Pop()
		for _, next := range(node.table) {
			iterator.stack.Push(next)
		}
	}
}