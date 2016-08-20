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

type Entry struct {
	key string
	value interface{}
}

func newEntry(key string, value interface{}) *Entry {
	return &Entry{
		key: key,
		value: value,
	}
}

func (entry *Entry) GetKey() string {
	return entry.key
}

func (entry *Entry) GetValue() interface{} {
	return entry.value
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

func (trie *Trie) KeySet() []string {
	keySet := []string{}
	keyTrace := []rune{}
	trie.keySetDfs(trie.root, &keySet, &keyTrace)
	return keySet
}

func (trie *Trie) keySetDfs(node *node, set *[]string, keyTrace *[]rune) {
	if node.hasValue() {
		*set = append(*set, string(*keyTrace))
	}
	for c, child := range(node.table) {
		*keyTrace = append(*keyTrace, c)
		trie.keySetDfs(child, set, keyTrace)
		*keyTrace = (*keyTrace)[:len(*keyTrace) - 1]
	}
}

func (trie *Trie) ValueSet() []interface{} {
	valueSet := []interface{}{}
	trie.valueSetDfs(trie.root, &valueSet)
	return valueSet
}

func (trie *Trie) valueSetDfs(node *node, set *[]interface{}) {
	if node.hasValue() {
		*set = append(*set, node.value)
	}
	for _, child := range(node.table) {
		trie.valueSetDfs(child, set)
	}
}

func (trie *Trie) EntrySet() []*Entry {
	entrySet := []*Entry{}
	keyTrace := []rune{}
	trie.entrySetDfs(trie.root, &entrySet, &keyTrace)
	return entrySet
}

func (trie *Trie) entrySetDfs(node *node, set *[]*Entry, keyTrace *[]rune) {
	if node.hasValue() {
		*set = append(*set, newEntry(string(*keyTrace), node.value))
	}
	for c, child := range(node.table) {
		*keyTrace = append(*keyTrace, c)
		trie.entrySetDfs(child, set, keyTrace)
		*keyTrace = (*keyTrace)[:len(*keyTrace) - 1]
	}
}

func (trie *Trie) Iterator() Iterator {
	return newTrieIterator(trie)
}

type trieIterator struct {
	entrySet []*Entry
	position int
}

func newTrieIterator(trie *Trie) *trieIterator {
	return &trieIterator{
		entrySet: trie.EntrySet(),
		position: 0,
	}
}

func (iterator *trieIterator) HasNext() bool {
	return len(iterator.entrySet) > iterator.position
}

func (iterator *trieIterator) Next() interface{} {
	entry := iterator.entrySet[iterator.position]
	iterator.position++
	return entry
}