package lbadd

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

// btree is the main structure.
//
// "order" invariants:
// - every node except root must contain at least order-1 keys
// - every node may contain at most (2*order)-1 keys
type btree struct {
	root  *node
	size  int
	order int
}

// newBtree creates a new instance of Btree
func newBtree() *btree {
	return &btree{
		root:  nil,
		size:  0,
		order: defaultOrder,
	}
}

func newBtreeOrder(order int) *btree {
	return &btree{
		root:  nil,
		size:  0,
		order: order,
	}
}

// get searches for a specific key in the btree,
// returning a pointer to the resulting entry
// and a boolean as to whether it exists in the tree
func (b *btree) get(k key) (result *entry, exists bool) {
	if b.root == nil || len(b.root.entries) == 0 {
		return nil, false
	}

	return b.getNode(b.root, k)
}

func (b *btree) getNode(node *node, k key) (result *entry, exists bool) {
	i, exists := b.search(node.entries, k)
	if exists {
		return node.entries[i], true
	}

	if i > len(node.children) {
		return nil, false
	}

	return b.getNode(node.children[i], k)
}

// insert takes a key and value, creats a new
// entry and inserts it in the tree according to the key
func (b *btree) insert(k key, v value) {
	if b.root == nil {
		b.size++
		b.root = &node{
			parent:   nil,
			entries:  []*entry{{k, v}},
			children: []*node{},
		}
		return
	}

	b.insertNode(b.root, &entry{k, v})
}

// insertNode takes a node and the entry to insert
func (b *btree) insertNode(node *node, entry *entry) (inserted bool) {
	// If the root node is already full, we need to split it
	if node == b.root && node.isFull(b.order) {
		b.root = node.split()
	}

	idx, exists := b.search(node.entries, entry.key)

	// The entry already exists, so it should be updated
	if exists {
		node.entries[idx] = entry
		return false
	}

	// If the node is a leaf node, add entry to the entries list
	// We can guarantee that we have room as it would otherwise have
	// been split.
	if node.isLeaf() {
		node.entries = append(node.entries, nil)
		copy(node.entries[idx+1:], node.entries[idx:])
		node.entries[idx] = entry
		b.size++
		return true
	}

	// The node is not a leaf, so we we need to check
	// if the appropriate child is already full,
	// and conditionally split it. Otherwise traverse
	// to that child.
	if node.children[idx].isFull(b.order) {
		node.children[idx] = node.children[idx].split()
	}

	return b.insertNode(node.children[idx], entry)
}

// remove tries to delete an entry from the tree, and
// returns true if the entry was removed, and false if
// the key was not found in the tree
func (b *btree) remove(k key) (removed bool) {
	if b.root == nil {
		return false
	}

	return b.removeNode(b.root, k)
}

func (b *btree) removeNode(node *node, k key) (removed bool) {
	idx, exists := b.search(node.entries, k)

	// If the key exists in a leaf node, we can simply remove
	// it outright
	if exists && node.isLeaf() {
		b.size--
		node.entries = append(node.entries[:idx], node.entries[idx+1:]...)
		return true
	}

	return true
}

// search takes a slice of entries and a key, and returns
// the position that the key would fit relative to all
// other entries' keys.
// e.g.
//       b.search([1, 2, 4], 3) => (2, false)
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

// isFull returns a bool indication whether the node
// already contains the maximum number of entries
// allowed for a given order
func (n *node) isFull(order int) bool {
	return len(n.entries) >= ((order * 2) - 1)
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
