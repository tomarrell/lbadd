package main

const defaultOrder = 3

// storage defines the interface to be implemented by
// the b-tree
type storage interface {
	get(k key)
	put(k key, v value)
	remove(k key)
}

type (
	key   int
	value interface{}
)

// node defines the stuct which contains keys (entries) and
// the child nodes of a particular node in the b-tree
type node struct {
	parent   *node
	entries  []*entry
	children []*node
}

// entry is a key/value pair that is stored in the b-tree
type entry struct {
	key   key
	value value
}

// btree is the main structure
type btree struct {
	root  *node
	size  int
	order int
}

func newBtree() *btree {
	return &btree{
		order: defaultOrder,
	}
}

func (b *btree) insert(k key, v value) {
	if b.root == nil {
		b.root = &node{
			parent:   nil,
			entries:  []*entry{{k, v}},
			children: []*node{},
		}
		return
	}

	b.insertEntry(b.root, &entry{k, v})
}

func (b *btree) insertEntry(node *node, entry *entry) {
	if len(node.children) == 0 {
		b.search(node.entries, entry.key)
	}
}

// TODO
func (b *btree) get(k key) *entry {
	return &entry{}
}

func (b *btree) search(entries []*entry, k key) (index int, exists bool) {
	var (
		low  = 0
		mid  = 0
		high = len(entries) - 1
	)

	for low <= high {
		mid = (high + low) / 2

		entryKey := entries[mid].key
		switch {
		case k > entryKey:
			low = mid + 1
		case k < entryKey:
			high = mid - 1
		case k == entryKey:
			return mid, true
		}
	}

	return low, false
}
