package main

type Trie struct {
	Val     interface{}
	SubTrie map[rune]*Trie
}

func NewTrie() *Trie {
	return new(Trie)
}

func (trie *Trie) Put(key string, value interface{}) {
	current := trie
	for _, r := range key {
		child := current.SubTrie[r]
		if child == nil {
			if current.SubTrie == nil {
				current.SubTrie = map[rune]*Trie{}
			}
			child = NewTrie()
			current.SubTrie[r] = child
		}
		current = child
	}
	current.Val = value
}

func (trie *Trie) Get(key string) interface{} {
	current := trie
	for _, r := range key {
		current = current.SubTrie[r]
		if current == nil {
			return nil
		}
	}
	return current.Val
}
