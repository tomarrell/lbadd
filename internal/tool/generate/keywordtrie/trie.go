package main

type trie struct {
	val interface{}
	sub map[rune]*trie
}

func newTrie() *trie {
	return new(trie)
}

func (t *trie) Put(key string, value interface{}) {
	current := t
	for _, r := range key {
		child := current.sub[r]
		if child == nil {
			if current.sub == nil {
				current.sub = map[rune]*trie{}
			}
			child = newTrie()
			current.sub[r] = child
		}
		current = child
	}
	current.val = value
}

func (t *trie) Get(key string) interface{} {
	current := t
	for _, r := range key {
		current = current.sub[r]
		if current == nil {
			return nil
		}
	}
	return current.val
}
