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
				entries: []*entry{{2, nil}},
				children: []*node{
					{entries: []*entry{{1, 1}}},
					{entries: []*entry{{2, 2}, {3, 3}}},
				},
			},
			key:            1,
			expectedExists: true,
		},
		{
			name: "entry one level deep right of root",
			root: &node{
				entries: []*entry{{2, nil}},
				children: []*node{
					{entries: []*entry{{1, 1}}},
					{entries: []*entry{{2, 2}, {3, 3}}},
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
					{entries: []*entry{{2, 2}, {3, 3}}},
				},
			},
			key:            4,
			expectedExists: false,
		},
		{
			name: "depth = 3 found",
			root: &node{
				entries: []*entry{{2, nil}},
				children: []*node{
					{entries: []*entry{{1, 1}}},
					{
						entries: []*entry{{3, nil}},
						children: []*node{
							{entries: []*entry{{2, 2}}},
							{entries: []*entry{{3, 3}, {4, 4}}},
						},
					},
				},
			},
			key:            4,
			expectedExists: true,
		},
		{
			name: "depth = 3 not found",
			root: &node{
				entries: []*entry{{2, nil}},
				children: []*node{
					{entries: []*entry{{1, 1}}},
					{
						entries: []*entry{{3, nil}},
						children: []*node{
							{entries: []*entry{{2, 2}}},
							{entries: []*entry{{3, 3}, {4, 4}}},
						},
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
				entries: []*entry{{3, nil}},
				children: []*node{
					{entries: []*entry{{1, 1}, {2, 2}}},
					{entries: []*entry{{3, 3}, {4, 4}, {5, 5}}},
				},
			},
		},
		{
			name:  "even entries node",
			input: &node{parent: parent, entries: []*entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}}},
			expected: &node{
				parent:  parent,
				entries: []*entry{{3, nil}},
				children: []*node{
					{entries: []*entry{{1, 1}, {2, 2}}},
					{entries: []*entry{{3, 3}, {4, 4}}},
				},
			},
		},
		{
			name:  "no parent",
			input: &node{entries: []*entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}},
			root:  true,
			expected: &node{
				entries: []*entry{{3, nil}},
				children: []*node{
					{entries: []*entry{{1, 1}, {2, 2}}},
					{entries: []*entry{{3, 3}, {4, 4}, {5, 5}}},
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
				entries:  []*entry{{1, nil}},
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
		root *node
		size int
	}
	type args struct {
		node  *node
		entry *entry
	}

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
					entries: []*entry{{2, nil}},
					children: []*node{
						{entries: []*entry{{1, 1}}},
						{entries: []*entry{{2, 2}}},
					},
				},
				&entry{2, 2},
			},
			wantSize:     2,
			wantInserted: false,
		},
		{
			name:   "entry exists one level down right unbalanced",
			fields: fields{size: 2},
			args: args{
				&node{
					entries: []*entry{{1, nil}},
					children: []*node{
						{},
						{entries: []*entry{{1, 1}, {2, 2}}},
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
					entries: []*entry{{2, nil}},
					children: []*node{
						{entries: []*entry{{1, 1}}},
						{entries: []*entry{{2, 2}}},
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
					entries: []*entry{{3, nil}},
					children: []*node{
						{entries: []*entry{{1, 1}}},
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
			fields: fields{size: 6},
			args: args{
				&node{
					entries: []*entry{{10, nil}},
					children: []*node{
						{entries: []*entry{{3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}}},
						{entries: []*entry{{10, 10}}},
					},
				},
				&entry{1, 1},
			},
			wantSize:     7,
			wantInserted: true,
		},
		{
			name:   "entry inserted one level down right, would more than overflow",
			fields: fields{size: 4},
			args: args{
				&node{
					entries: []*entry{{10, nil}},
					children: []*node{
						{entries: []*entry{{3, 3}, {4, 4}, {5, 5}}},
						{entries: []*entry{{10, 10}, {11, 11}, {12, 12}, {13, 13}, {14, 14}, {15, 15}, {16, 16}, {17, 17}, {18, 18}, {19, 19}, {29, 29}}},
					},
				},
				&entry{30, 30},
			},
			wantSize:     5,
			wantInserted: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &btree{
				root:  tt.fields.root,
				size:  tt.fields.size,
				order: 3,
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
			fields:      fields{root: nil, size: 0},
			args:        args{k: 1},
			wantRemoved: false,
			wantSize:    0,
		},
		{
			name: "remove entry from root",
			fields: fields{
				size: 1,
				root: &node{
					entries: []*entry{{1, 1}},
				},
			},
			args:        args{k: 1},
			wantRemoved: true,
			wantSize:    0,
		},
		{
			name: "entry doesn't exist",
			fields: fields{
				size:  2,
				order: 2,
				root: &node{
					entries: []*entry{{2, 2}},
					children: []*node{
						{entries: []*entry{{1, 1}}},
						{},
					},
				},
			},
			args:        args{k: 3},
			wantRemoved: false,
			wantSize:    2,
		},
		{
			name: "entry multiple levels down",
			fields: fields{
				size:  21,
				order: 2,
				root: &node{
					entries: []*entry{{30, nil}, {60, nil}},
					children: []*node{
						{
							entries: []*entry{{10, nil}, {20, nil}},
							children: []*node{
								{entries: []*entry{{2, 2}, {4, 4}}},
								{entries: []*entry{{12, 12}, {18, 18}}},
								{entries: []*entry{{21, 21}, {28, 28}, {29, 29}}},
							},
						},
						{
							entries: []*entry{{40, nil}, {50, nil}},
							children: []*node{
								{entries: []*entry{{34, 34}, {36, 36}, {37, 37}}},
								{entries: []*entry{{41, 41}, {43, 43}}},
								{entries: []*entry{{58, 58}, {59, 59}}},
							},
						},
						{
							entries: []*entry{{70, nil}, {80, nil}},
							children: []*node{
								{entries: []*entry{{65, 65}, {68, 68}, {69, 69}}},
								{entries: []*entry{{70, 70}, {71, 71}}},
								{entries: []*entry{{80, 80}, {100, 100}}},
							},
						},
					},
				},
			},
			args:        args{k: 36},
			wantRemoved: true,
			wantSize:    20,
		},
		{
			name: "entry multi level doesn't exist",
			fields: fields{
				size:  21,
				order: 2,
				root: &node{
					entries: []*entry{{30, nil}, {60, nil}},
					children: []*node{
						{
							entries: []*entry{{10, nil}, {20, nil}},
							children: []*node{
								{entries: []*entry{{2, 2}, {4, 4}}},
								{entries: []*entry{{12, 12}, {18, 18}}},
								{entries: []*entry{{21, 21}, {28, 28}, {29, 29}}},
							},
						},
						{
							entries: []*entry{{40, nil}, {50, nil}},
							children: []*node{
								{entries: []*entry{{34, 34}, {36, 36}, {37, 37}}},
								{entries: []*entry{{41, 41}, {43, 43}}},
								{entries: []*entry{{58, 58}, {59, 59}}},
							},
						},
						{
							entries: []*entry{{70, nil}, {80, nil}},
							children: []*node{
								{entries: []*entry{{65, 65}, {68, 68}, {69, 69}}},
								{entries: []*entry{{70, 70}, {71, 71}}},
								{entries: []*entry{{80, 80}, {100, 100}}},
							},
						},
					},
				},
			},
			args:        args{k: 99},
			wantRemoved: false,
			wantSize:    21,
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

func TestRemove_structure(t *testing.T) {
	order := 3
	tree := func() *btree {
		return &btree{
			order: order,
			size:  8,
			root: &node{
				parent:  nil,
				entries: []*entry{{5, nil}, {10, nil}},
				children: []*node{
					{entries: []*entry{{1, 1}, {2, 2}, {3, 3}}},
					{entries: []*entry{{5, 5}, {7, 7}}},
					{entries: []*entry{{10, 10}, {20, 20}, {21, 21}}},
				},
			},
		}
	}

	tests := []struct {
		name      string
		haveTree  *btree
		removeKey key
		wantTree  *btree
	}{
		{
			name:      "remove entry from left leaf, no underflow",
			haveTree:  tree(),
			removeKey: 3,
			wantTree: &btree{
				order: order,
				size:  7,
				root: &node{
					entries: []*entry{{5, nil}, {10, nil}},
					children: []*node{
						{entries: []*entry{{1, 1}, {2, 2}}},
						{entries: []*entry{{5, 5}, {7, 7}}},
						{entries: []*entry{{10, 10}, {20, 20}, {21, 21}}},
					},
				},
			},
		},
		{
			name:      "entry doesn't exist",
			haveTree:  tree(),
			removeKey: 9,
			wantTree: &btree{
				order: order,
				size:  8,
				root: &node{
					entries: []*entry{{5, nil}, {10, nil}},
					children: []*node{
						{entries: []*entry{{1, 1}, {2, 2}, {3, 3}}},
						{entries: []*entry{{5, 5}, {7, 7}}},
						{entries: []*entry{{10, 10}, {20, 20}, {21, 21}}},
					},
				},
			},
		},
		// {
		// name:      "remove entry from middle leaf, trigger underflow",
		// haveTree:  tree(),
		// removeKey: 7,
		// wantTree: &btree{
		// order: order,
		// size:  7,
		// root: &node{
		// entries: []*entry{{5, nil}, {10, nil}},
		// children: []*node{
		// {entries: []*entry{{1, 1}, {2, 2}, {3, 3}}},
		// {entries: []*entry{{5, 5}, {7, 7}}},
		// {entries: []*entry{{10, 10}, {20, 20}, {21, 21}}},
		// },
		// },
		// },
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.haveTree.remove(tt.removeKey)
			assert.Equal(t, tt.wantTree, tt.haveTree)
		})
	}
}

func TestNode_isFull(t *testing.T) {
	type fields struct {
		parent   *node
		entries  []*entry
		children []*node
	}
	type args struct {
		order int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "order 3, node not full",
			fields: fields{entries: []*entry{}},
			args:   args{order: 3},
			want:   false,
		},
		{
			name:   "order 3, node full",
			fields: fields{entries: []*entry{{1, 1}, {2, 2}, {3, 3}}},
			args:   args{order: 3},
			want:   true,
		},
		{
			name:   "order 3, node nearly full",
			fields: fields{entries: []*entry{{1, 1}, {2, 2}}},
			args:   args{order: 3},
			want:   false,
		},
		{
			name:   "order 3, node over filled (bug case)",
			fields: fields{entries: []*entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}}},
			args:   args{order: 3},
			want:   true,
		},
		{
			name:   "order 5, node full",
			fields: fields{entries: []*entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}},
			args:   args{order: 5},
			want:   true,
		},
		{
			name:   "order 5, node almost full",
			fields: fields{entries: []*entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}}},
			args:   args{order: 5},
			want:   false,
		},
		{
			name:   "order 5, node empty",
			fields: fields{entries: []*entry{}},
			args:   args{order: 5},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				parent:   tt.fields.parent,
				entries:  tt.fields.entries,
				children: tt.fields.children,
			}

			got := n.isFull(tt.args.order)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNode_canSteal(t *testing.T) {
	type fields struct {
		parent   *node
		entries  []*entry
		children []*node
	}
	type args struct {
		order int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				parent:   tt.fields.parent,
				entries:  tt.fields.entries,
				children: tt.fields.children,
			}

			got := n.canSteal(tt.args.order)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_btree_getAll(t *testing.T) {
	type fields struct {
		root *node
		size int
	}
	type args struct {
		limit int
	}

	root := &node{}
	root.entries = []*entry{{4, 4}, {8, 8}}
	root.children = []*node{
		{
			parent:  root,
			entries: []*entry{{0, 0}, {1, 1}, {2, 2}},
		},
		{
			parent:  root,
			entries: []*entry{{5, 5}, {7, 7}},
		},
		{
			parent:  root,
			entries: []*entry{{9, 9}, {11, 11}, {12, 12}},
		},
	}

	f := fields{root: root, size: 10}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*entry
	}{
		{
			name:   "empty tree returns empty entries",
			fields: fields{size: 0},
			args:   args{limit: 10},
			want:   []*entry{},
		},
		{
			name:   "limit of 0 returns empty entries",
			fields: f,
			args:   args{limit: 0},
			want:   []*entry{},
		},
		{
			name:   "returns all entries up to limit",
			fields: f,
			args:   args{limit: 0},
			want:   []*entry{},
		},
		// TODO add more cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &btree{
				root: tt.fields.root,
				size: tt.fields.size,
			}

			got := b.getAll(tt.args.limit)
			assert.Equal(t, tt.want, got)
		})
	}
}
