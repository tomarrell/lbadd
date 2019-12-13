package main

type btree struct {
	root  *node
	size  int
	order int
}

type node struct {
	parent   *node
	entries  []*entry
	children []*node
}

type entry struct {
	key   interface{}
	value interface{}
}

func newBtree(order int) *btree {
	return &btree{
		order: order,
	}
}
