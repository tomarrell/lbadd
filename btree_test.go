package lbadd

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

	order := 3

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bt := newBtreeOrder(order)

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

func TestGet(t *testing.T) {
	cases := []struct {
		name           string
		root           *node
		key            key
		expectedExists bool
	}{
		{
			name:           "no root",
			root:           nil,
			expectedExists: false,
		},
		{
			name:           "empty root",
			root:           &node{},
			expectedExists: false,
		},
		{
			name:           "entries only in root",
			root:           &node{entries: []*entry{{1, 1}, {2, 2}, {3, 3}}},
			key:            2,
			expectedExists: true,
		},
		{
			name: "entry one level deep left of root",
			root: &node{
				entries: []*entry{{2, 2}},
				children: []*node{
					{entries: []*entry{{1, 1}}},
					{entries: []*entry{{3, 3}}},
				},
			},
			key:            1,
			expectedExists: true,
		},
		{
			name: "entry one level deep right of root",
			root: &node{
				entries: []*entry{{2, 2}},
				children: []*node{
					{entries: []*entry{{1, 1}}},
					{entries: []*entry{{3, 3}}},
				},
			},
			key:            3,
			expectedExists: true,
		},
		{
			name: "depth > 1 and key not exist",
			root: &node{
				entries: []*entry{{2, 2}},
				children: []*node{
					{entries: []*entry{{1, 1}}},
					{entries: []*entry{{3, 3}}},
				},
			},
			key:            4,
			expectedExists: false,
		},
		{
			name: "depth = 3 found",
			root: &node{
				entries: []*entry{{2, 2}},
				children: []*node{
					{entries: []*entry{{1, 1}}},
					{
						entries:  []*entry{{3, 3}},
						children: []*node{{}, {entries: []*entry{{4, 4}}}},
					},
				},
			},
			key:            4,
			expectedExists: true,
		},
		{
			name: "depth = 3 not found",
			root: &node{
				entries: []*entry{{2, 2}},
				children: []*node{
					{entries: []*entry{{1, 1}}},
					{
						entries:  []*entry{{3, 3}},
						children: []*node{{}, {entries: []*entry{{4, 4}}}},
					},
				},
			},
			key:            5,
			expectedExists: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			btree := newBtree()
			btree.root = tc.root

			_, exists := btree.get(tc.key)
			assert.Equal(t, tc.expectedExists, exists)
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
			name:    "single value, already exists",
			entries: []*entry{{key: 1}},
			key:     1,
			exists:  true,
			index:   0,
		},
		{
			name:    "already exists",
			entries: []*entry{{key: 1}, {key: 2}, {key: 4}, {key: 5}},
			key:     4,
			exists:  true,
			index:   2,
		},
		{
			name:    "doc example",
			entries: []*entry{{key: 1}, {key: 2}, {key: 4}},
			key:     3,
			exists:  false,
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

func TestInsertNode(t *testing.T) {
	type fields struct {
		root  *node
		size  int
		order int
	}
	type args struct {
		node  *node
		entry *entry
	}

	order := 3

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantSize     int
		wantInserted bool
	}{
		{
			name:         "insert single entry",
			fields:       fields{},
			args:         args{&node{}, &entry{1, 1}},
			wantSize:     1,
			wantInserted: true,
		},
		{
			name:         "entry already exists",
			fields:       fields{size: 1},
			args:         args{&node{entries: []*entry{{1, 1}}}, &entry{1, 1}},
			wantSize:     1,
			wantInserted: false,
		},
		{
			name:   "entry exists one level down right",
			fields: fields{size: 2},
			args: args{
				&node{
					entries: []*entry{{1, 1}},
					children: []*node{
						{},
						{entries: []*entry{{2, 2}}},
					},
				},
				&entry{2, 2},
			},
			wantSize:     2,
			wantInserted: false,
		},
		{
			name:   "entry exists one level down left",
			fields: fields{size: 2},
			args: args{
				&node{
					entries: []*entry{{2, 2}},
					children: []*node{
						{entries: []*entry{{1, 1}}},
						{},
					},
				},
				&entry{1, 1},
			},
			wantSize:     2,
			wantInserted: false,
		},
		{
			name:   "entry inserted one level down right",
			fields: fields{size: 2},
			args: args{
				&node{
					entries: []*entry{{2, 2}},
					children: []*node{
						{},
						{entries: []*entry{{3, 3}}},
					},
				},
				&entry{4, 4},
			},
			wantSize:     3,
			wantInserted: true,
		},
		{
			name:   "entry inserted one level down left, would overflow",
			fields: fields{size: 2},
			args: args{
				&node{
					entries: []*entry{{10, 10}},
					children: []*node{
						{entries: []*entry{{3, 3}, {4, 4}, {5, 5}}},
						{},
					},
				},
				&entry{1, 1},
			},
			wantSize:     3,
			wantInserted: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &btree{
				root:  tt.fields.root,
				size:  tt.fields.size,
				order: order,
			}

			got := b.insertNode(tt.args.node, tt.args.entry)
			assert.Equal(t, tt.wantInserted, got)
			assert.Equal(t, tt.wantSize, b.size)
		})
	}
}

func TestRemove(t *testing.T) {
	type fields struct {
		root  *node
		size  int
		order int
	}
	type args struct {
		k key
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		wantRemoved bool
		wantSize    int
	}{
		{
			name:        "no root",
			fields:      fields{root: nil, size: 0, order: 3},
			args:        args{k: 1},
			wantRemoved: false,
			wantSize:    0,
		},
		{
			name:        "remove entry from root",
			fields:      fields{root: &node{entries: []*entry{{1, 1}}}, size: 1, order: 3},
			args:        args{k: 1},
			wantRemoved: true,
			wantSize:    0,
		},
		{
			name: "remove entry depth 1 left",
			fields: fields{
				size:  2,
				order: 3,
				root: &node{
					entries: []*entry{{1, 1}},
					children: []*node{
						nil,
						{entries: []*entry{{2, 2}}},
					},
				},
			},
			args:        args{k: 2},
			wantRemoved: true,
			wantSize:    1,
		},
		{
			name: "remove entry depth 1 right",
			fields: fields{
				size:  2,
				order: 3,
				root: &node{
					entries: []*entry{{2, 2}},
					children: []*node{
						{entries: []*entry{{1, 1}}},
					},
				},
			},
			args:        args{k: 1},
			wantRemoved: true,
			wantSize:    1,
		},
		{
			name: "key doesn't exist",
			fields: fields{
				size:  2,
				order: 3,
				root: &node{
					entries: []*entry{{2, 2}},
					children: []*node{
						{entries: []*entry{{1, 1}}},
					},
				},
			},
			args:        args{k: 3},
			wantRemoved: false,
			wantSize:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &btree{
				root:  tt.fields.root,
				size:  tt.fields.size,
				order: tt.fields.order,
			}

			gotRemoved := b.remove(tt.args.k)
			assert.Equal(t, tt.wantRemoved, gotRemoved)
			assert.Equal(t, tt.wantSize, b.size)
		})
	}
}
