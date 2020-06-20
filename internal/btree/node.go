package btree

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
// are three options:
//   A: steal from the right sibling
//   B: steal from the left sibling
//   C: merge the node with the left sibling
//   D: merge the node with the right sibling
//
func (n *node) recursiveBalance(k Key, order int, b *Btree) {
	// If we're at the root or the current node is not breaking the invariant, we
	// can stop.
	if n.isRoot() || !n.isUnderflowed(order) {
		return
	}

	idx, _ := search(n.entries, k)

	// Scenario A: steal from the right sibling
	//
	// First we check whether a right sibling exists. If it does and it contains
	// enough entries for us to steal one, we proceed.
	//
	// We will steal right right sibling's leftmost child and append this to the
	// current node's children.
	//
	if rSib, exists := n.rightSibling(k); exists && rSib.canStealEntry(order) {
		// Entry operations
		//
		// Append the right sibling's leftmost entry to this node
		n.entries = append(n.entries, rSib.entries[0])
		// Remove the right sibling's leftmost entry
		rSib.entries = rSib.entries[1:]

		// Child operations
		// TODO update the stolen child's parent
		if !n.isLeaf() {
			// Append the right sibling's leftmost child to this node
			n.children = append(n.children, rSib.children[0])
			// Remove the right sibling's leftmost child
			rSib.children = rSib.children[1:]
		}

		// Parent operations
		//
		parIdx, _ := search(n.parent.entries, k)

		// Replace the parent key to the right sibling's leftmost entry's key
		n.parent.entries[parIdx] = &Entry{rSib.entries[0].key, nil}
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
		// Entry operations
		//
		// Prepend the left sibling's rightmost entry to this node
		n.entries = append([]*Entry{lSib.entries[len(lSib.entries)-1]}, n.entries...)
		// Remove the left sibling's rightmost entry
		lSib.entries = lSib.entries[:len(lSib.entries)-1]

		// Child operations
		// TODO update the stolen child's parent
		if !n.isLeaf() {
			// Prepend the left sibling's rightmost child to this node
			toSteal := lSib.children[len(lSib.children)-1]
			n.children = append([]*node{toSteal}, n.children...)
			// Remove the left sibling's rightmost child
			lSib.children = lSib.children[:len(lSib.children)-1]
		}

		parIdx, _ := search(n.parent.entries, k)

		// Set the parent's key to the key of the last entry in the sibling
		n.parent.entries[parIdx] = &Entry{lSib.entries[len(lSib.entries)-1].key, nil}
		// Remove the left sibling's last entry
		lSib.entries = lSib.entries[:len(lSib.entries)-1]

		return
	}

	// Try to merge left
	if _, exists := n.leftSibling(k); exists {
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

	for !n.isRoot() && par != nil {
		count++
		par = par.parent
	}

	return count
}
