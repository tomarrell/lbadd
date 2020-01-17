// Btree contains the btree struct, which is used as the primary data store of
// the database. It is an implementation of a traditional B+tree, however will
// from here on out just be referred to as "btree".
//
// The btree supports 3 primary operations:
// - get: given a key, retrieve the corresponding entry
// - put: given a key and a value, create an entry in the btree
// - remove: given a key, remove the corresponding entry in the tree if it
// exists

package lbadd

import (
	"fmt"
	"math"
)

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
	i, exists := search(node.entries, k)
	if exists && node.isLeaf() {
		return node.entries[i-1], true
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
	idx, exists := search(node.entries, entry.key)

	// The entry already exists, so it should be updated
	if exists {
		node.entries[idx-1] = entry
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
	idx, exists := search(node.entries, k)

	// If the node is not a leaf, we need to continue traversal
	if !node.isLeaf() {
		// If it exists, the idx is one less than what we need
		fmt.Println("traversing to on child index", idx)
		return b.removeNode(node.children[idx], k)
	}

	// Otherwise, we check if the entry exists, and return if it doesn't
	if !exists {
		return false
	}

	// Ok, so we've found the key, now we need to remove it.
	node.entries = append(node.entries[:idx-1], node.entries[idx:]...)
	b.size--

	// Now we need to check if we've caused an underflow
	if node.isUnderflowed(b.order) {
		// Can steal from the left leaf sibling
		lleaf, exists := node.leftSibling(k)
		if exists && lleaf.canSteal(b.order) {
			panic("can steal from left sibling")
		}

		// Can steal from the right leaf sibling
		rleaf, exists := node.rightSibling(k)
		if exists && rleaf.canSteal(b.order) {
			// Append the right sibling's first entry to this node
			node.entries = append(node.entries, rleaf.entries[0])
			// Remove the right sibling's first entry
			rleaf.entries = rleaf.entries[1:]
			// Replace the parent key to the right sibling's first entry's key
			node.parent.entries[idx-1] = &entry{rleaf.entries[0].key, nil}
			return true
		}

		// Can't steal from either left or right, so we're going to have to merge
		fmt.Println("found nothing to steal")

		// Try to merge left
		_, exists = node.leftSibling(k)
		if exists {
			panic("can merge into left sibling")
		}

		dest, exists := node.rightSibling(k)
		if exists {
			fmt.Println(dest)
			panic("can merge into right sibling")
		}

		return false
	} else {
		fmt.Println("no underflow")
	}

	return true
}

//
func (b *btree) getAll(limit int) []*entry {
	if b.size == 0 || limit == 0 {
		return []*entry{}
	}

	panic("unimplemented")
}

//
func (b *btree) getAbove(k key, limit int) []*entry {
	panic("unimplemented")
}

//
func (b *btree) getBelow(k key, limit int) []*entry {
	panic("unimplemented")
}

//
func (b *btree) getBetween(low, high key, limit int) []*entry {
	panic("unimplemented")
}

// search takes a slice of entries and a key, and returns
// the position that the key would fit relative to all
// other entries' keys.
// e.g.
//       b.search([1, 2, 4], 3) => (2, false)
func search(entries []*entry, k key) (index int, exists bool) {
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

func (n *node) isLeaf() bool {
	return len(n.children) == 0
}

// isFull returns a bool indication whether the node
// already contains the maximum number of entries
// allowed for a given order
func (n *node) isFull(order int) bool {
	return len(n.entries) >= order
}

// canSteal returns a bool indicating whether or not
// the node contains enough entries to be able to take one
func (n *node) canSteal(order int) bool {
	return len(n.entries) > int(math.Ceil(float64(order)/2.0))
}

// Returns true when the node has too few entries to
// satisfy the order invariant, given a specific order
func (n *node) isUnderflowed(order int) bool {
	return len(n.entries) < int(math.Ceil(float64(order)/2.0))
}

// returns whether the node can successfully be split into
// two children while maintaining the invariants
func (n *node) canSplit(order int) bool {
	return float64(len(n.entries)) >= 2*math.Ceil(float64(order)/2.0)
}

// leftSibling returns the left sibling if it exists, indicating such
func (n *node) leftSibling(k key) (sibling *node, exists bool) {
	parIdx, exists := search(n.parent.entries, k)
	if parIdx == 0 {
		return nil, false
	}

	return n.parent.children[parIdx-1], true
}

// rightSibling returns the left sibling if it exists, indicating such
func (n *node) rightSibling(k key) (sibling *node, exists bool) {
	parIdx, _ := search(n.parent.entries, k)

	if parIdx+1 >= len(n.parent.children) {
		return nil, false
	}

	return n.parent.children[parIdx+1], true
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

	return &node{
		parent:   n.parent,
		entries:  []*entry{{n.entries[mid].key, nil}},
		children: append(n.children, left, right),
	}
}
