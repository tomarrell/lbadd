package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTree(t *testing.T) {
	t.Skip()
	cases := []struct {
		name   string
		insert []entry
		get    []entry
	}{
		{
			name:   "set and get",
			insert: []entry{{1, 1}},
			get:    []entry{{1, 1}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bt := newBtree()

			for _, e := range tc.insert {
				bt.insert(e.key, e.value)
			}

			for _, g := range tc.get {
				assert.Equal(t,
					g.value,
					func() *entry {
						e, _ := bt.get(g.key)
						return e
					}(),
				)
			}
		})
	}
}

func TestKeySearch(t *testing.T) {
	cases := []struct {
		name    string
		entries []*entry
		key     key
		exists  bool
		index   int
	}{
		{
			name:    "single value",
			entries: []*entry{{key: 1}},
			key:     2,
			exists:  false,
			index:   1,
		},
		{
			name:    "single value duplicate",
			entries: []*entry{{key: 1}},
			key:     1,
			exists:  true,
			index:   0,
		},
		{
			name:    "duplicate",
			entries: []*entry{{key: 1}, {key: 2}, {key: 4}, {key: 5}},
			key:     4,
			exists:  true,
			index:   2,
		},
		{
			name:    "no entries",
			entries: []*entry{},
			key:     2,
			exists:  false,
			index:   0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bt := newBtree()

			idx, exists := bt.search(tc.entries, tc.key)

			assert.Equal(t, tc.exists, exists)
			assert.Equal(t, tc.index, idx)
		})
	}
}

func TestNodeSplit(t *testing.T) {
	parent := &node{}

	cases := []struct {
		name     string
		root     bool
		input    *node
		expected *node
	}{
		{
			name:  "simple node",
			input: &node{parent: parent, entries: []*entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}},
			expected: &node{
				parent:  parent,
				entries: []*entry{{3, 3}},
				children: []*node{
					{entries: []*entry{{1, 1}, {2, 2}}},
					{entries: []*entry{{4, 4}, {5, 5}}},
				},
			},
		},
		{
			name:  "even entries node",
			input: &node{parent: parent, entries: []*entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}}},
			expected: &node{
				parent:  parent,
				entries: []*entry{{3, 3}},
				children: []*node{
					{entries: []*entry{{1, 1}, {2, 2}}},
					{entries: []*entry{{4, 4}}},
				},
			},
		},
		{
			name:  "no parent",
			input: &node{entries: []*entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}},
			root:  true,
			expected: &node{
				entries: []*entry{{3, 3}},
				children: []*node{
					{entries: []*entry{{1, 1}, {2, 2}}},
					{entries: []*entry{{4, 4}, {5, 5}}},
				},
			},
		},
		{
			name:  "empty node",
			input: &node{parent: parent, entries: []*entry{}},
			expected: &node{
				parent:   parent,
				entries:  []*entry{},
				children: []*node{},
			},
		},
		{
			name:  "single entry",
			input: &node{parent: parent, entries: []*entry{{1, 1}}},
			expected: &node{
				parent:   parent,
				entries:  []*entry{{1, 1}},
				children: []*node{},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			newNode := tc.input.split()
			assert.Equal(t, tc.expected.parent, newNode.parent)
			assert.Equal(t, tc.expected.entries, newNode.entries)

			for i := range tc.expected.children {
				expectedChild, newChild := tc.expected.children[i], newNode.children[i]

				assert.Equal(t, &tc.input, &newChild.parent)
				assert.Equal(t, expectedChild.entries, newChild.entries)
				assert.Equal(t, expectedChild.children, newChild.children)
			}
		})
	}
}
