// Btree contains the btree struct, which is used as the primary data store of
// the database. It is an implementation of a traditional B+tree, however will
// from here on out just be referred to as "btree".
//
// The btree supports 3 primary operations:
// - get: given a key, retrieve the corresponding entry
// - put: given a key and a value, create an entry in the btree
// - remove: given a key, remove the corresponding entry in the tree if it
//   exists

package btree

const defaultOrder = 3

// btree is an implementation of a B+tree with the following invariants
//
// ref: c = len(children), k = len(keys), o = order
//
// - all leaves must be same distance (d) from root
// - the root node has at least two children
// - every node must have k+1 references
// - every internal node has at least ceil(o / 2) children
// - every leaf node contains at least ceil(o / 2) keys
// - for every internal node N with k: all keys in the first child's subtree are
//   less than N's first key; and all keys in the i'th child's subtree (2 ≤ i ≤ k)
//   are between the (i − 1)th key of n and the i'th key of n
type Btree struct {
	root  *node
	size  int
	order int
}

// newBtree creates a new instance of Btree
func NewBtree() *Btree {
	return &Btree{
		root:  nil,
		size:  0,
		order: defaultOrder,
	}
}

func NewBtreeOrder(order int) *Btree {
	return &Btree{
		root:  nil,
		size:  0,
		order: order,
	}
}

// Output the btree as a string representation
func (b *Btree) String() string {
	out := ""

	out += "R: "
	out += b.root.String()
	out += "\n"

	queue := b.root.children
	// Breadth traversal
	for len(queue) > 0 {
		// pop from front of queue and append the popped nodes children
		n := queue[0]
		queue = append(queue[1:], n.children...)

		out += n.String()
	}

	return out
}

// get searches for a specific Key in the btree, returning a pointer to the
// resulting entry and a boolean as to whether it exists in the tree
func (b *Btree) Get(k Key) (result *Entry, exists bool) {
	if b.root == nil || len(b.root.entries) == 0 {
		return nil, false
	}

	return b.getNode(b.root, k)
}

func (b *Btree) getNode(node *node, k Key) (result *Entry, exists bool) {
	i, exists := search(node.entries, k)
	if exists && node.isLeaf() {
		return node.entries[i-1], true
	}

	if i > len(node.children) {
		return nil, false
	}

	return b.getNode(node.children[i], k)
}

// insert takes a Key and value, creats a new entry and inserts it in the tree
// according to the Key
func (b *Btree) Insert(k Key, v Value) {
	if b.root == nil {
		b.size++
		b.root = &node{
			parent:   nil,
			entries:  []*Entry{{k, v}},
			children: []*node{},
		}
		return
	}

	b.insertNode(b.root, &Entry{k, v})
}

// insertNode takes a node and the entry to insert
func (b *Btree) insertNode(node *node, entry *Entry) (inserted bool) {
	// If the root node is already full, we need to split it
	if node == b.root && node.isFull(b.order) {
		b.root = node.split()
	}

	// Search for the Key in the node's entries
	idx, exists := search(node.entries, entry.key)

	// The entry already exists, so it should be updated
	if exists {
		node.entries[idx-1] = entry
		return false
	}

	// If the node is a leaf node, add entry to the entries list We can guarantee
	// that we have room as it would otherwise have been split.
	if node.isLeaf() {
		node.entries = append(node.entries, nil)
		copy(node.entries[idx+1:], node.entries[idx:])
		node.entries[idx] = entry
		b.size++
		return true
	}

	// The node is not a leaf, so we we need to check if the appropriate child is
	// already full, and conditionally split it. Otherwise traverse to that child.
	if node.children[idx].isFull(b.order) {
		node.children[idx] = node.children[idx].split()
	}

	return b.insertNode(node.children[idx], entry)
}

// remove tries to delete an entry from the tree, and returns true if the entry
// was removed, and false if the Key was not found in the tree
func (b *Btree) Remove(k Key) (removed bool) {
	if b.root == nil {
		return false
	}

	return b.removeNode(b.root, k)
}

// removeNode takes a node and Key and bool, and recursively deletes k from the
// node, while maintaining the order invariants
func (b *Btree) removeNode(n *node, k Key) (removed bool) {
	idx, exists := search(n.entries, k)

	// If the node is not a leaf, we need to continue traversal
	if !n.isLeaf() {
		// If it exists, the idx is one less than what we need
		return b.removeNode(n.children[idx], k)
	}

	// Otherwise, we check if the entry exists, and return if it doesn't
	if !exists {
		return false
	}

	// Ok, so we've found the Key, now we need to remove it.
	n.entries = append(n.entries[:idx-1], n.entries[idx:]...)
	b.size--

	// Now we need to check if we've caused an underflow
	if n.isUnderflowed(b.order) {
		n.recursiveBalance(k, b.order, b)

		return true
	}

	return true
}

//
func (b *Btree) GetAll(limit int) []*Entry {
	if b.size == 0 || limit == 0 {
		return []*Entry{}
	}

	panic("unimplemented")
}

//
func (b *Btree) GetAbove(k Key, limit int) []*Entry {
	panic("unimplemented")
}

//
func (b *Btree) GetBelow(k Key, limit int) []*Entry {
	panic("unimplemented")
}

//
func (b *Btree) GetBetween(low, high Key, limit int) []*Entry {
	panic("unimplemented")
}

// search takes a slice of entries and a Key, and returns the position that the
// Key would fit relative to all other entries' Keys.
// e.g.
//       b.search([1, 2, 4], 3) => (2, false)
func search(entries []*Entry, k Key) (index int, exists bool) {
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
			return mid + 1, true
		}
	}

	return low, false
}
