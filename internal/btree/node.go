package btree

import "github.com/davecgh/go-spew/spew"

// node defines the stuct which contains keys (entries) and the child nodes of a
// particular node in the b-tree
type node struct {
	parent   *node
	entries  []*Entry
	children []*node
}

// recursiveBalance ...
//
// We know that a node is below the threshold for which we must rebalance. There
// are four options:
//   A: steal from the right sibling
//   B: steal from the left sibling
//   C: merge the node with the left sibling
//   D: merge the node with the right sibling
//
// TODO it probably makes sense to have this rebalance method on the *btree
// rather than the node. However, it is convenient here for recursion.
func (n *node) recursiveBalance(k Key, order int, b *Btree) {
	// If we're at the root or the current node is not breaking the invariant, we
	// can stop.
	if n.isRoot() || !n.isUnderflowed(order) {
		return
	}

	// Scenario A: steal from the right sibling
	//
	// First we check whether a right sibling exists. If it does and it contains
	// enough entries for us to steal one, we proceed.
	//
	// We will steal right right sibling's leftmost child and append this to the
	// current node's children.
	//
	if rSib, exists := n.rightSibling(k); exists && rSib.canStealEntry(order) {
		spew.Println("Stealing from right sibling")
		// Entry operations
		//
		// Append the right sibling's leftmost entry to this node
		n.entries = append(n.entries, rSib.entries[0])
		// Remove the right sibling's leftmost entry
		rSib.entries = rSib.entries[1:]

		// Child operations
		if n.isInternal() {
			// Append the right sibling's leftmost child to this node
			n.children = append(n.children, rSib.children[0])
			// Update the stolen child's parent
			n.children[len(n.children)-1].parent = n
			// Remove the right sibling's leftmost child
			rSib.children = rSib.children[1:]
		}

		// Parent operations
		//
		parIdx, _ := search(n.parent.entries, k)

		// Replace the parent key to the right sibling's leftmost entry's key
		n.parent.entries[parIdx] = &Entry{rSib.entries[0].key, nil}

		spew.Printf("New entries: %s\n", n.entries)
		spew.Printf("New parent: %s\n", n.parent.entries)
		spew.Printf("New right sibling: %s\n", rSib.entries)

		return
	}

	// Scenario B: steal from left sibling.
	//
	// First we check whether a left sibling exists. If it does and it contains
	// enough entries for us to steal one, we proceed.
	//
	// We will steal the left sibling's right most entry, and we will need to
	// steal the right most child as well. This will be prepended to the current
	// node's children.
	//
	// We will then update the entry key in the parent to the key of the entry
	// that we stole from the left sibling.
	//
	if lSib, exists := n.leftSibling(k); exists && lSib.canStealEntry(order) {
		spew.Println("Stealing from left sibling")

		// Entry operations
		//
		// Prepend the left sibling's rightmost entry to this node
		n.entries = append([]*Entry{lSib.entries[len(lSib.entries)-1]}, n.entries...)
		// Remove the left sibling's rightmost entry
		lSib.entries = lSib.entries[:len(lSib.entries)-1]
		// Rotate the parent's key down

		// Child operations
		//
		if n.isInternal() {
			// Prepend the left sibling's rightmost child to this node
			n.children = append([]*node{lSib.children[len(lSib.children)-1]}, n.children...)
			n.children[0].parent = n
			// Remove the left sibling's rightmost child
			lSib.children = lSib.children[:len(lSib.children)-1]
		}

		// Parent operations
		//
		parIdx, _ := search(n.parent.entries, k)

		if n.isInternal() {
			// Rotate the parent down
			stolenKey := n.entries[0].key
			n.entries[0].key = n.parent.entries[parIdx-1].key
			n.parent.entries[parIdx-1].key = stolenKey
		}

		// Set the parent's key to the key of the first entry of the node
		n.parent.entries[parIdx-1] = &Entry{n.entries[0].key, nil}

		// spew.Printf("New entries: %s\n", n.entries)
		// spew.Printf("New parent: %s\n", n.parent.entries)
		// spew.Printf("New right sibling: %s\n", rSib.entries)

		return
	}

	// Scenario C: merge the node with the left sibling.
	//
	// First we check whether a left sibling exists.
	//
	// TODO handle the children
	//
	if _, exists := n.leftSibling(k); exists {
		spew.Println("Merging left sibling")
		panic("can merge left")
	}

	// Scenario D: merge the node with the right sibling.
	//
	// First we check whether a right sibling exists. If it does, we know that it
	// doesn't have enough to steal. Therefore we must merge the current node and
	// it. We can do this as we know both nodes together will not break the
	// maximum number of children invariant.
	//
	// In order to merge right, we append all of the right nodes' entries to the
	// current node.
	//
	// We then need to remove the parent at index i, and update the parent at
	// index i-1 to have the key of the leftmost entry.
	//
	// Once this is done, we must now check if the parent breaks and of the
	// invariants. If it does, we should recursively call this method on the
	// parent.
	//
	if rSib, exists := n.rightSibling(k); exists {
		spew.Println("Merging right sibling")
		parIdx, _ := search(n.parent.entries, k)

		// Add the right sibling's entries to the current node
		n.entries = append(n.entries, rSib.entries...)
		n.children = append(n.children, rSib.children...)

		// Remove the right sibling from the parent
		n.parent.children = append(n.parent.children[:parIdx+1], n.parent.children[parIdx+2:]...)
		// Remove the right key entry from parent
		n.parent.entries = append(n.parent.entries[:parIdx], n.parent.entries[parIdx+1:]...)

		if n.parent.isUnderflowed(b.order) {
			n.parent.recursiveBalance(k, order, b)
			return
		}
	}
}

func (n *node) isLeaf() bool {
	return len(n.children) == 0
}

func (n *node) isInternal() bool {
	return !n.isLeaf()
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

	for !n.isRoot() && par != nil {
		count++
		par = par.parent
	}

	return count
}
