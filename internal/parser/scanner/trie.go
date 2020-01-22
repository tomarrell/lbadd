package scanner

import "strings"

// Trier exposes the Trie structure capabilities.
type Trier interface {
	Get(key string) interface{}
	Put(key string, value interface{}) bool
	Delete(key string) bool
	Walk(walker WalkFunc) error
}

// WalkFunc defines some action to take on the given key and value during
// a Trie Walk. Returning a non-nil error will terminate the Walk.
type WalkFunc func(key string, value interface{}) error

// StringSegmenter takes a string key with a starting index and returns
// the first segment after the start and the ending index. When the end is
// reached, the returned nextIndex should be -1.
// Implementations should NOT allocate heap memory as Trie Segmenters are
// called upon Gets. See PathSegmenter.
type StringSegmenter func(key string, start int) (segment string, nextIndex int)

// PathSegmenter segments string key paths by slash separators. For example,
// "/a/b/c" -> ("/a", 2), ("/b", 4), ("/c", -1) in successive calls. It does
// not allocate any heap memory.
func PathSegmenter(path string, start int) (segment string, next int) {
	if len(path) == 0 || start < 0 || start > len(path)-1 {
		return "", -1
	}
	end := strings.IndexRune(path[start+1:], '/') // next '/' after 0th rune
	if end == -1 {
		return path[start:], -1
	}
	return path[start : start+end+1], start + end + 1
}

// RuneTrie is a trie of runes with string keys and interface{} values.
// Note that internal nodes have nil values so a stored nil value will not
// be distinguishable and will not be included in Walks.
type RuneTrie struct {
	value    interface{}
	children map[rune]*RuneTrie
}

// NewRuneTrie allocates and returns a new *RuneTrie.
func NewRuneTrie() *RuneTrie {
	return &RuneTrie{
		children: make(map[rune]*RuneTrie),
	}
}

// Get returns the value stored at the given key. Returns nil for internal
// nodes or for nodes with a value of nil.
func (trie *RuneTrie) Get(key string) interface{} {
	node := trie
	itr := -1
	for _, r := range key {
		node = node.children[r]
		itr++
		if node == nil {
			return itr
		}
	}
	return node.value
}

// Put inserts the value into the trie at the given key, replacing any
// existing items. It returns true if the put adds a new value, false
// if it replaces an existing value.
// Note that internal nodes have nil values so a stored nil value will not
// be distinguishable and will not be included in Walks.
func (trie *RuneTrie) Put(key string, value interface{}) bool {
	node := trie
	for _, r := range key {
		child, _ := node.children[r]
		if child == nil {
			child = NewRuneTrie()
			node.children[r] = child
		}
		node = child
	}
	// does node have an existing value?
	isNewVal := node.value == nil
	node.value = value
	return isNewVal
}

// Delete removes the value associated with the given key. Returns true if a
// node was found for the given key. If the node or any of its ancestors
// becomes childless as a result, it is removed from the trie.
func (trie *RuneTrie) Delete(key string) bool {
	path := make([]nodeRune, len(key)) // record ancestors to check later
	node := trie
	for i, r := range key {
		path[i] = nodeRune{r: r, node: node}
		node = node.children[r]
		if node == nil {
			// node does not exist
			return false
		}
	}
	// delete the node value
	node.value = nil
	// if leaf, remove it from its parent's children map. Repeat for ancestor path.
	if node.isLeaf() {
		// iterate backwards over path
		for i := len(key) - 1; i >= 0; i-- {
			parent := path[i].node
			r := path[i].r
			delete(parent.children, r)
			if parent.value != nil || !parent.isLeaf() {
				// parent has a value or has other children, stop
				break
			}
		}
	}
	return true // node (internal or not) existed and its value was nil'd
}

// Walk iterates over each key/value stored in the trie and calls the given
// walker function with the key and value. If the walker function returns
// an error, the walk is aborted.
// The traversal is depth first with no guaranteed order.
func (trie *RuneTrie) Walk(walker WalkFunc) error {
	return trie.walk("", walker)
}

// RuneTrie node and the rune key of the child the path descends into.
type nodeRune struct {
	node *RuneTrie
	r    rune
}

func (trie *RuneTrie) walk(key string, walker WalkFunc) error {
	if trie.value != nil {
		if err := walker(key, trie.value); err != nil {
			return err
		}
	}
	for r, child := range trie.children {
		if err := child.walk(key+string(r), walker); err != nil {
			return err
		}
	}
	return nil
}

func (trie *RuneTrie) isLeaf() bool {
	return len(trie.children) == 0
}
