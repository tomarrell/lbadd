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
		root:  nil,
		size:  0,
		order: defaultOrder,
	}
}

func (b *btree) get(k key) (result *entry, exists bool) {
	if b.root == nil {
		return nil, false
	}

	return b.getNode(b.root, k)
}

func (b *btree) getNode(node *node, k key) (result *entry, exists bool) {
	return nil, false
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

	b.insertNode(b.root, &entry{k, v})
}

// TODO finish this
func (b *btree) insertNode(node *node, entry *entry) (inserted bool) {
	idx, exists := b.search(node.entries, entry.key)
	if exists {
		node.entries[idx] = entry
		return false
	}

	// If the root node would be filled, we need to split it
	if node == b.root && node.wouldFill(b.order) {
		b.root = node.split()
	}

	// If the node is a leaf node, put it into the entries list
	if node.isLeaf() {
		node.entries = append(node.entries, nil)
		copy(node.entries[idx+1:], node.entries[idx:])
		node.entries[idx] = entry
		return true
	}

	return b.insertNode(node.children[idx], entry)
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

func (n *node) isLeaf() bool {
	return len(n.children) == 0
}

func (n *node) wouldFill(order int) bool {
	return len(n.entries)+1 >= ((order * 2) - 1)
}

// Splits a full node to have a single, median,
// entry, and two child nodes containing the left
// and right halves of the entries
func (n *node) split() *node {
	if len(n.entries) == 0 {
		return n
	}

	mid := len(n.entries) / 2

	left := &node{
		parent:  n,
		entries: append([]*entry{}, n.entries[:mid]...),
	}
	right := &node{
		parent:  n,
		entries: append([]*entry{}, n.entries[mid+1:]...),
	}

	n.entries = []*entry{n.entries[mid]}
	n.children = append(n.children, left, right)

	return n
}
