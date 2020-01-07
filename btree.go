// Btree contains the btree struct, which is used as the primary data store of
// the database.
//
// The btree supports 3 primary operations:
// - get: given a key, retrieve the corresponding entry
// - put: given a key and a value, create an entry in the btree
// - remove: given a key, remove the corresponding entry in the tree if it
// exists

package lbadd

const defaultOrder = 3

// storage defines the interface to be implemented by
// the b-tree
type storage interface {
	get(k key) (v *entry, exists bool)
	insert(k key, v value)
	remove(k key) (removed bool)
	getAll(limit int) []*entry
	getAbove(k key, limit int) []*entry
	getBelow(k key, limit int) []*entry
	getBetween(low, high key, limit int) []*entry
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

	// Search for the key in the node's entries
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

// removeNode takes a node and key and bool, and recursively deletes
// k from the node, while maintaining the order invariants
func (b *btree) removeNode(node *node, k key) (removed bool) {
	idx, exists := b.search(node.entries, k)

	// If the key exists in a leaf node, we can simply remove
	// it outright
	if node.isLeaf() {
		if exists {
			b.size--
			node.entries = append(node.entries[:idx], node.entries[idx+1:]...)
			return true
		}
		// We've reached the bottom and couldn't find the key
		return false
	}

	// If the key exists in the node, but it is not a leaf
	if exists {
		child := node.children[idx]
		// There are enough entries in left child to take one
		if child.canSteal(b.order) {
			stolen := child.entries[len(child.entries)-1]
			node.entries[idx] = stolen
			return b.removeNode(child, stolen.key)
		}

		// child = node.children[idx]
		// There are enough entries in the right child to take one
		// if child.canSteal(b.order) {
		// TODO implement this
		// }

		// Both children don't have enough entries, so we need
		// to merge the left and right children and take a key
		// TODO
		return false
	}

	return b.removeNode(node.children[idx], k)
}

//
func (b *btree) getAll(limit int) []*entry {
	if b.size == 0 || limit == 0 {
		return []*entry{}
	}

	// TODO unimplemented

	return nil
}

//
func (b *btree) getAbove(k key, limit int) []*entry {
	// TODO unimplemented
	return []*entry{}
}

//
func (b *btree) getBelow(k key, limit int) []*entry {
	// TODO unimplemented
	return []*entry{}
}

//
func (b *btree) getBetween(low, high key, limit int) []*entry {
	// TODO unimplemented
	return []*entry{}
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

// canSteal returns a bool indicating whether or not
// the node contains enough entries to be able to take one
func (n *node) canSteal(order int) bool {
	return len(n.entries)-1 > order-1
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
		entries: append([]*entry{}, n.entries[mid:]...),
	}

	n.entries = []*entry{{n.entries[mid].key, nil}}
	n.children = append(n.children, left, right)

	return n
}
