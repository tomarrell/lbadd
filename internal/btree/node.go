package btree

import (
	"fmt"
)

// node defines the stuct which contains keys (entries) and the child nodes of a
// particular node in the b-tree
type node struct {
	parent   *node
	entries  []*Entry
	children []*node
}

func (n *node) String() string {
	nStr := ""

	nStr += fmt.Sprintf("%v", n.entries)

	return nStr
}

func (n *node) recursiveBalance(k Key, order int, b *Btree) {
	if n.isRoot() {
		return
	}

	idx, exists := search(n.entries, k)
	// Can steal from the left leaf sibling
	lleaf, exists := n.leftSibling(k)
	if exists && lleaf.canStealEntry(order) {
		// Take the parent's key and prepend it to this node
		n.entries = append([]*Entry{n.parent.entries[idx]}, n.entries...)
		// Set the parent's key to the key of the last entry in the sibling
		n.parent.entries[idx] = &Entry{lleaf.entries[len(lleaf.entries)-1].key, nil}
		// Remove the left sibling's last entry
		lleaf.entries = lleaf.entries[:len(lleaf.entries)-1]

		// If the node isn't a leaf, we need to steal the sibling's rightmost child
		// as well, making sure it knows who its new parent is
		if !n.isLeaf() {
			stolen := lleaf.children[len(lleaf.children)-1]
			stolen.parent = n
			n.children = append([]*node{stolen}, n.children...)
			lleaf.children = lleaf.children[:len(lleaf.children)-1]
		}

		return
	}

	// Can steal from the right leaf sibling
	rleaf, exists := n.rightSibling(k)
	if exists && rleaf.canStealEntry(order) {
		// Append the right sibling's first entry to this node
		n.entries = append(n.entries, rleaf.entries[0])
		// Remove the right sibling's first entry
		rleaf.entries = rleaf.entries[1:]
		// Replace the parent key to the right sibling's first entry's key
		n.parent.entries[idx] = &Entry{rleaf.entries[0].key, nil}
		return
	}

	// Try to merge left
	_, exists = n.leftSibling(k)
	if exists {
		panic("can merge left")
	}

	// Try to merge right
	rdest, exists := n.rightSibling(k)
	if exists {
		// Create a new node, which is the combination of the two nodes to merge
		mergedNode := &node{
			parent:   rdest.parent,
			entries:  append(n.entries, rdest.entries...),
			children: append(n.children, rdest.children...),
		}

		// Check if removing a child from the parent would cause it to underflow
		parIdx, _ := search(n.parent.entries, k)

		n.parent.entries = append(n.parent.entries[:parIdx], n.parent.entries[parIdx+1:]...)
		n.parent.children = append(n.parent.children[:parIdx], n.parent.children[parIdx+1:]...)
		n.parent.children[parIdx] = mergedNode

		// Check if parent now has too few children, and recursively merge
		if n.parent.isUnderflowed(order) {
			n.parent.recursiveBalance(k, order, b)
			return
		}

		panic("can merge right")
	}
}

func (n *node) isLeaf() bool {
	return len(n.children) == 0
}

// isRoot returns whether or not the current node is the root of the tree
func (n *node) isRoot() bool {
	return n.parent == nil
}

// isFull returns a bool indication whether the node already contains the
// maximum number of entries allowed for a given order
func (n *node) isFull(order int) bool {
	return len(n.entries) >= order
}

// canSteal returns a bool indicating whether or not the node contains enough
// entries to be able to take one
func (n *node) canStealEntry(order int) bool {
	if n.isLeaf() {
		return len(n.entries) > order/2
	}

	return len(n.children) > order/2
}

// Returns true when the node has too few entries to satisfy the order
// invariant, given a specific order
func (n *node) isUnderflowed(order int) bool {
	if n.isLeaf() {
		return len(n.entries) < order/2
	}

	return len(n.children) < order/2
}

// returns whether the node can successfully be split into two children while
// maintaining the invariants
func (n *node) canSplit(order int) bool {
	return len(n.children) >= 2*order
}

// leftSibling returns the left sibling if it exists, indicating such
func (n *node) leftSibling(k Key) (sibling *node, exists bool) {
	parIdx, exists := search(n.parent.entries, k)
	if parIdx == 0 {
		return nil, false
	}

	return n.parent.children[parIdx-1], true
}

// rightSibling returns the left sibling if it exists, indicating such
func (n *node) rightSibling(k Key) (sibling *node, exists bool) {
	parIdx, _ := search(n.parent.entries, k)

	if parIdx+1 >= len(n.parent.children) {
		return nil, false
	}

	return n.parent.children[parIdx+1], true
}

// Splits a full node to have a single, median, entry, and two child nodes
// containing the left and right halves of the entries
func (n *node) split() *node {
	if len(n.entries) == 0 {
		return n
	}

	mid := len(n.entries) / 2

	left := &node{
		parent:  n,
		entries: append([]*Entry{}, n.entries[:mid]...),
	}
	right := &node{
		parent:  n,
		entries: append([]*Entry{}, n.entries[mid:]...),
	}

	return &node{
		parent:   n.parent,
		entries:  []*Entry{{n.entries[mid].key, nil}},
		children: append(n.children, left, right),
	}
}

// depth returns the depth of the current node from the root
func (n *node) depth() int {
	count := 0
	par := n.parent

	for !par.isRoot() {
		count++
		par = par.parent
	}

	return count
}
