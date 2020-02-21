package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBtree(t *testing.T) {
	t.Skip()
	cases := []struct {
		name   string
		insert []Entry
		get    []Entry
	}{
		{
			name:   "set and get",
			insert: []Entry{{1, 1}},
			get:    []Entry{{1, 1}},
		},
	}

	order := 3

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bt := NewBtreeOrder(order)

			for _, e := range tc.insert {
				bt.Insert(e.key, e.value)
			}

			for _, g := range tc.get {
				assert.Equal(t,
					g.value,
					func() *Entry {
						e, _ := bt.Get(g.key)
						return e
					}(),
				)
			}
		})
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		name       string
		root       *node
		key        Key
		wantEntry  *Entry
		wantExists bool
	}{
		{
			name:       "no root",
			root:       nil,
			wantExists: false,
		},
		{
			name:       "empty root",
			root:       &node{},
			wantExists: false,
		},
		{
			name:       "entries only in root",
			root:       &node{entries: []*Entry{{1, 1}, {2, 2}, {3, 3}}},
			key:        2,
			wantEntry:  &Entry{2, 2},
			wantExists: true,
		},
		{
			name: "Entry one level deep left of root",
			root: &node{
				entries: []*Entry{{2, nil}},
				children: []*node{
					{entries: []*Entry{{1, 1}}},
					{entries: []*Entry{{2, 2}, {3, 3}}},
				},
			},
			key:        1,
			wantEntry:  &Entry{1, 1},
			wantExists: true,
		},
		{
			name: "Entry one level deep right of root",
			root: &node{
				entries: []*Entry{{2, nil}},
				children: []*node{
					{entries: []*Entry{{1, 1}}},
					{entries: []*Entry{{2, 2}, {3, 3}}},
				},
			},
			key:        3,
			wantEntry:  &Entry{3, 3},
			wantExists: true,
		},
		{
			name: "depth > 1 and Key not exist",
			root: &node{
				entries: []*Entry{{2, 2}},
				children: []*node{
					{entries: []*Entry{{1, 1}}},
					{entries: []*Entry{{2, 2}, {3, 3}}},
				},
			},
			key:        4,
			wantExists: false,
		},
		{
			name: "depth = 3 found",
			root: &node{
				entries: []*Entry{{2, nil}},
				children: []*node{
					{entries: []*Entry{{1, 1}}},
					{
						entries: []*Entry{{3, nil}},
						children: []*node{
							{entries: []*Entry{{2, 2}}},
							{entries: []*Entry{{3, 3}, {4, 4}}},
						},
					},
				},
			},
			key:        4,
			wantEntry:  &Entry{4, 4},
			wantExists: true,
		},
		{
			name: "depth = 3 not found",
			root: &node{
				entries: []*Entry{{2, nil}},
				children: []*node{
					{entries: []*Entry{{1, 1}}},
					{
						entries: []*Entry{{3, nil}},
						children: []*node{
							{entries: []*Entry{{2, 2}}},
							{entries: []*Entry{{3, 3}, {4, 4}}},
						},
					},
				},
			},
			key:        5,
			wantExists: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := NewBtree()
			b.root = tc.root

			Entry, exists := b.Get(tc.key)
			assert.Equal(t, tc.wantEntry, Entry)
			assert.Equal(t, tc.wantExists, exists)
		})
	}
}

func TestKeySearch(t *testing.T) {
	cases := []struct {
		name    string
		entries []*Entry
		key     Key
		exists  bool
		index   int
	}{
		{
			name:    "single value",
			entries: []*Entry{{key: 1}},
			key:     2,
			exists:  false,
			index:   1,
		},
		{
			name:    "single value, already exists",
			entries: []*Entry{{key: 1}},
			key:     1,
			exists:  true,
			index:   1,
		},
		{
			name:    "already exists",
			entries: []*Entry{{key: 1}, {key: 2}, {key: 4}, {key: 5}},
			key:     4,
			exists:  true,
			index:   3,
		},
		{
			name:    "doc example",
			entries: []*Entry{{key: 1}, {key: 2}, {key: 4}},
			key:     3,
			exists:  false,
			index:   2,
		},
		{
			name:    "no entries",
			entries: []*Entry{},
			key:     2,
			exists:  false,
			index:   0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			idx, exists := search(tc.entries, tc.key)

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
			input: &node{parent: parent, entries: []*Entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}},
			expected: &node{
				parent:  parent,
				entries: []*Entry{{3, nil}},
				children: []*node{
					{entries: []*Entry{{1, 1}, {2, 2}}},
					{entries: []*Entry{{3, 3}, {4, 4}, {5, 5}}},
				},
			},
		},
		{
			name:  "even entries node",
			input: &node{parent: parent, entries: []*Entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}}},
			expected: &node{
				parent:  parent,
				entries: []*Entry{{3, nil}},
				children: []*node{
					{entries: []*Entry{{1, 1}, {2, 2}}},
					{entries: []*Entry{{3, 3}, {4, 4}}},
				},
			},
		},
		{
			name:  "no parent",
			input: &node{entries: []*Entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}},
			root:  true,
			expected: &node{
				entries: []*Entry{{3, nil}},
				children: []*node{
					{entries: []*Entry{{1, 1}, {2, 2}}},
					{entries: []*Entry{{3, 3}, {4, 4}, {5, 5}}},
				},
			},
		},
		{
			name:  "empty node",
			input: &node{parent: parent, entries: []*Entry{}},
			expected: &node{
				parent:   parent,
				entries:  []*Entry{},
				children: []*node{},
			},
		},
		{
			name:  "single Entry",
			input: &node{parent: parent, entries: []*Entry{{1, 1}}},
			expected: &node{
				parent:   parent,
				entries:  []*Entry{{1, nil}},
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
		Entry *Entry
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantSize     int
		wantInserted bool
	}{
		{
			name:         "insert single Entry",
			fields:       fields{},
			args:         args{&node{}, &Entry{1, 1}},
			wantSize:     1,
			wantInserted: true,
		},
		{
			name:         "Entry already exists",
			fields:       fields{size: 1},
			args:         args{&node{entries: []*Entry{{1, 1}}}, &Entry{1, 1}},
			wantSize:     1,
			wantInserted: false,
		},
		{
			name:   "Entry exists one level down right",
			fields: fields{size: 2},
			args: args{
				&node{
					entries: []*Entry{{2, nil}},
					children: []*node{
						{entries: []*Entry{{1, 1}}},
						{entries: []*Entry{{2, 2}}},
					},
				},
				&Entry{2, 2},
			},
			wantSize:     2,
			wantInserted: false,
		},
		{
			name:   "Entry exists one level down right unbalanced",
			fields: fields{size: 2},
			args: args{
				&node{
					entries: []*Entry{{1, nil}},
					children: []*node{
						{},
						{entries: []*Entry{{1, 1}, {2, 2}}},
					},
				},
				&Entry{2, 2},
			},
			wantSize:     2,
			wantInserted: false,
		},
		{
			name:   "Entry exists one level down left",
			fields: fields{size: 2},
			args: args{
				&node{
					entries: []*Entry{{2, nil}},
					children: []*node{
						{entries: []*Entry{{1, 1}}},
						{entries: []*Entry{{2, 2}}},
					},
				},
				&Entry{1, 1},
			},
			wantSize:     2,
			wantInserted: false,
		},
		{
			name:   "Entry inserted one level down right",
			fields: fields{size: 2},
			args: args{
				&node{
					entries: []*Entry{{3, nil}},
					children: []*node{
						{entries: []*Entry{{1, 1}}},
						{entries: []*Entry{{3, 3}}},
					},
				},
				&Entry{4, 4},
			},
			wantSize:     3,
			wantInserted: true,
		},
		{
			name:   "Entry inserted one level down left, would overflow",
			fields: fields{size: 6},
			args: args{
				&node{
					entries: []*Entry{{10, nil}},
					children: []*node{
						{entries: []*Entry{{3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}}},
						{entries: []*Entry{{10, 10}}},
					},
				},
				&Entry{1, 1},
			},
			wantSize:     7,
			wantInserted: true,
		},
		{
			name:   "Entry inserted one level down right, would more than overflow",
			fields: fields{size: 4},
			args: args{
				&node{
					entries: []*Entry{{10, nil}},
					children: []*node{
						{entries: []*Entry{{3, 3}, {4, 4}, {5, 5}}},
						{entries: []*Entry{{10, 10}, {11, 11}, {12, 12}, {13, 13}, {14, 14}, {15, 15}, {16, 16}, {17, 17}, {18, 18}, {19, 19}, {29, 29}}},
					},
				},
				&Entry{30, 30},
			},
			wantSize:     5,
			wantInserted: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Btree{
				root:  tt.fields.root,
				size:  tt.fields.size,
				order: 3,
			}

			got := b.insertNode(tt.args.node, tt.args.Entry)
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
		k Key
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
			name: "remove Entry from root",
			fields: fields{
				size: 1,
				root: &node{
					entries: []*Entry{{1, 1}},
				},
			},
			args:        args{k: 1},
			wantRemoved: true,
			wantSize:    0,
		},
		{
			name: "Entry doesn't exist",
			fields: fields{
				size:  2,
				order: 2,
				root: &node{
					entries: []*Entry{{2, 2}},
					children: []*node{
						{entries: []*Entry{{1, 1}}},
						{},
					},
				},
			},
			args:        args{k: 3},
			wantRemoved: false,
			wantSize:    2,
		},
		{
			name: "Entry multiple levels down",
			fields: fields{
				size:  21,
				order: 2,
				root: &node{
					entries: []*Entry{{30, nil}, {60, nil}},
					children: []*node{
						{
							entries: []*Entry{{10, nil}, {20, nil}},
							children: []*node{
								{entries: []*Entry{{2, 2}, {4, 4}}},
								{entries: []*Entry{{12, 12}, {18, 18}}},
								{entries: []*Entry{{21, 21}, {28, 28}, {29, 29}}},
							},
						},
						{
							entries: []*Entry{{40, nil}, {50, nil}},
							children: []*node{
								{entries: []*Entry{{34, 34}, {36, 36}, {37, 37}}},
								{entries: []*Entry{{41, 41}, {43, 43}}},
								{entries: []*Entry{{58, 58}, {59, 59}}},
							},
						},
						{
							entries: []*Entry{{70, nil}, {80, nil}},
							children: []*node{
								{entries: []*Entry{{65, 65}, {68, 68}, {69, 69}}},
								{entries: []*Entry{{70, 70}, {71, 71}}},
								{entries: []*Entry{{80, 80}, {100, 100}}},
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
			name: "Entry multi level doesn't exist",
			fields: fields{
				size:  21,
				order: 2,
				root: &node{
					entries: []*Entry{{30, nil}, {60, nil}},
					children: []*node{
						{
							entries: []*Entry{{10, nil}, {20, nil}},
							children: []*node{
								{entries: []*Entry{{2, 2}, {4, 4}}},
								{entries: []*Entry{{12, 12}, {18, 18}}},
								{entries: []*Entry{{21, 21}, {28, 28}, {29, 29}}},
							},
						},
						{
							entries: []*Entry{{40, nil}, {50, nil}},
							children: []*node{
								{entries: []*Entry{{34, 34}, {36, 36}, {37, 37}}},
								{entries: []*Entry{{41, 41}, {43, 43}}},
								{entries: []*Entry{{58, 58}, {59, 59}}},
							},
						},
						{
							entries: []*Entry{{70, nil}, {80, nil}},
							children: []*node{
								{entries: []*Entry{{65, 65}, {68, 68}, {69, 69}}},
								{entries: []*Entry{{70, 70}, {71, 71}}},
								{entries: []*Entry{{80, 80}, {100, 100}}},
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
			b := &Btree{
				root:  tt.fields.root,
				size:  tt.fields.size,
				order: tt.fields.order,
			}

			gotRemoved := b.Remove(tt.args.k)
			assert.Equal(t, tt.wantRemoved, gotRemoved)
			assert.Equal(t, tt.wantSize, b.size)
		})
	}
}

func TestRemove_structure(t *testing.T) {
	t.Skip()
	order := 3
	tree := func() *Btree {
		return &Btree{
			order: order,
			size:  8,
			root: repairParents(t, &node{
				parent:  nil,
				entries: []*Entry{{5, nil}, {10, nil}},
				children: []*node{
					{entries: []*Entry{{1, 1}, {2, 2}, {3, 3}}},
					{entries: []*Entry{{5, 5}, {7, 7}}},
					{entries: []*Entry{{10, 10}, {20, 20}, {21, 21}}},
				},
			}, nil),
		}
	}

	tests := []struct {
		name     string
		haveTree *Btree
		remove   Key
		wantTree *Btree
	}{
		{
			name:     "remove Entry from left leaf, no underflow",
			haveTree: tree(),
			remove:   3,
			wantTree: &Btree{
				order: order,
				size:  7,
				root: repairParents(t, &node{
					entries: []*Entry{{5, nil}, {10, nil}},
					children: []*node{
						{entries: []*Entry{{1, 1}, {2, 2}}},
						{entries: []*Entry{{5, 5}, {7, 7}}},
						{entries: []*Entry{{10, 10}, {20, 20}, {21, 21}}},
					},
				}, nil),
			},
		},
		{
			name:     "Entry doesn't exist",
			haveTree: tree(),
			remove:   9,
			wantTree: &Btree{
				order: order,
				size:  8,
				root: repairParents(t, &node{
					entries: []*Entry{{5, nil}, {10, nil}},
					children: []*node{
						{entries: []*Entry{{1, 1}, {2, 2}, {3, 3}}},
						{entries: []*Entry{{5, 5}, {7, 7}}},
						{entries: []*Entry{{10, 10}, {20, 20}, {21, 21}}},
					},
				}, nil),
			},
		},
		{
			name:     "remove Entry from middle leaf, trigger underflow, have enough to combine",
			haveTree: tree(),
			remove:   7,
			wantTree: &Btree{
				order: order,
				size:  7,
				root: repairParents(t, &node{
					entries: []*Entry{{3, nil}, {10, nil}},
					children: []*node{
						{entries: []*Entry{{1, 1}, {2, 2}}},
						{entries: []*Entry{{3, 3}, {5, 5}}},
						{entries: []*Entry{{10, 10}, {20, 20}, {21, 21}}},
					},
				}, nil),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.haveTree.Remove(tt.remove)

			assert.Equal(t, tt.wantTree, tt.haveTree)
		})
	}
}

func TestRemove_structure_2(t *testing.T) {
	order := 4
	tree := func() *Btree {
		return &Btree{
			order: order,
			size:  11,
			root: repairParents(t, &node{
				entries: []*Entry{{13, nil}},
				children: []*node{
					{
						entries: []*Entry{{9, nil}, {11, nil}},
						children: []*node{
							{entries: []*Entry{{1, 1}, {4, 4}}},
							{entries: []*Entry{{9, 9}, {10, 10}}},
							{entries: []*Entry{{11, 11}, {12, 12}}},
						},
					},
					{
						entries: []*Entry{{16, nil}},
						children: []*node{
							{entries: []*Entry{{13, 13}, {15, 15}}},
							{entries: []*Entry{{16, 16}, {20, 20}, {25, 25}}},
						},
					},
				},
			}, nil),
		}
	}

	tests := []struct {
		name     string
		haveTree *Btree
		remove   []Key
		wantTree *Btree
	}{
		{
			name:     "remove 13",
			haveTree: tree(),
			remove:   []Key{13},
			wantTree: &Btree{
				order: order,
				size:  10,
				root: repairParents(t, &node{
					entries: []*Entry{{13, nil}},
					children: []*node{
						{
							entries: []*Entry{{9, nil}, {11, nil}},
							children: []*node{
								{entries: []*Entry{{1, 1}, {4, 4}}},
								{entries: []*Entry{{9, 9}, {10, 10}}},
								{entries: []*Entry{{11, 11}, {12, 12}}},
							},
						},
						{
							entries: []*Entry{{20, nil}},
							children: []*node{
								{entries: []*Entry{{15, 15}, {16, 16}}},
								{entries: []*Entry{{20, 20}, {25, 25}}},
							},
						},
					},
				}, nil),
			},
		},
		{
			name:     "remove 13, 15",
			haveTree: tree(),
			remove:   []Key{13, 15},
			wantTree: &Btree{
				order: order,
				size:  9,
				root: repairParents(t, &node{
					entries: []*Entry{{11, nil}},
					children: []*node{
						{
							entries: []*Entry{{9, nil}},
							children: []*node{
								{entries: []*Entry{{1, 1}, {4, 4}}},
								{entries: []*Entry{{9, 9}, {10, 10}}},
							},
						},
						{
							entries: []*Entry{{13, nil}},
							children: []*node{
								{entries: []*Entry{{11, 11}, {12, 12}}},
								{entries: []*Entry{{16, 16}, {20, 20}, {25, 25}}},
							},
						},
					},
				}, nil),
			},
		},
		{
			name:     "remove 13, 15, 1",
			haveTree: tree(),
			remove:   []Key{13, 15, 1},
			wantTree: &Btree{
				order: order,
				size:  8,
				root: repairParents(t, &node{
					entries: []*Entry{{11, nil}, {13, nil}},
					children: []*node{
						{
							entries: []*Entry{{4, 4}, {9, 9}, {10, 10}},
						},
						{
							entries: []*Entry{{11, 11}, {12, 12}},
						},
						{
							entries: []*Entry{{16, 16}, {20, 20}, {25, 25}},
						},
					},
				}, nil),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, k := range tt.remove {
				assert.True(t, tt.haveTree.Remove(k))
			}

			assert.Equal(t, tt.wantTree, tt.haveTree)
		})
	}
}

func TestNode_isFull(t *testing.T) {
	type fields struct {
		parent   *node
		entries  []*Entry
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
			fields: fields{entries: []*Entry{}},
			args:   args{order: 3},
			want:   false,
		},
		{
			name:   "order 3, node full",
			fields: fields{entries: []*Entry{{1, 1}, {2, 2}, {3, 3}}},
			args:   args{order: 3},
			want:   true,
		},
		{
			name:   "order 3, node nearly full",
			fields: fields{entries: []*Entry{{1, 1}, {2, 2}}},
			args:   args{order: 3},
			want:   false,
		},
		{
			name:   "order 3, node over filled (bug case)",
			fields: fields{entries: []*Entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}}},
			args:   args{order: 3},
			want:   true,
		},
		{
			name:   "order 5, node full",
			fields: fields{entries: []*Entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}},
			args:   args{order: 5},
			want:   true,
		},
		{
			name:   "order 5, node almost full",
			fields: fields{entries: []*Entry{{1, 1}, {2, 2}, {3, 3}, {4, 4}}},
			args:   args{order: 5},
			want:   false,
		},
		{
			name:   "order 5, node empty",
			fields: fields{entries: []*Entry{}},
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

func TestNode_canStealEntry(t *testing.T) {
	type fields struct {
		parent   *node
		entries  []*Entry
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

			got := n.canStealEntry(tt.args.order)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_Btree_getAll(t *testing.T) {
	type fields struct {
		root *node
		size int
	}
	type args struct {
		limit int
	}

	root := &node{}
	root.entries = []*Entry{{4, 4}, {8, 8}}
	root.children = []*node{
		{
			parent:  root,
			entries: []*Entry{{0, 0}, {1, 1}, {2, 2}},
		},
		{
			parent:  root,
			entries: []*Entry{{5, 5}, {7, 7}},
		},
		{
			parent:  root,
			entries: []*Entry{{9, 9}, {11, 11}, {12, 12}},
		},
	}

	f := fields{root: root, size: 10}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Entry
	}{
		{
			name:   "empty tree returns empty entries",
			fields: fields{size: 0},
			args:   args{limit: 10},
			want:   []*Entry{},
		},
		{
			name:   "limit of 0 returns empty entries",
			fields: f,
			args:   args{limit: 0},
			want:   []*Entry{},
		},
		{
			name:   "returns all entries up to limit",
			fields: f,
			args:   args{limit: 0},
			want:   []*Entry{},
		},
		// TODO add more cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Btree{
				root: tt.fields.root,
				size: tt.fields.size,
			}

			got := b.GetAll(tt.args.limit)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_node_canSplit(t *testing.T) {
	type fields struct {
		parent   *node
		entries  []*Entry
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				parent:   tt.fields.parent,
				entries:  tt.fields.entries,
				children: tt.fields.children,
			}
			if got := n.canSplit(tt.args.order); got != tt.want {
				t.Errorf("node.canSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_isUnderflowed(t *testing.T) {
	type fields struct {
		parent   *node
		entries  []*Entry
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				parent:   tt.fields.parent,
				entries:  tt.fields.entries,
				children: tt.fields.children,
			}
			if got := n.isUnderflowed(tt.args.order); got != tt.want {
				t.Errorf("node.isUnderflowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_isLeaf(t *testing.T) {
	type fields struct {
		parent   *node
		entries  []*Entry
		children []*node
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				parent:   tt.fields.parent,
				entries:  tt.fields.entries,
				children: tt.fields.children,
			}
			if got := n.isLeaf(); got != tt.want {
				t.Errorf("node.isLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_leftSibling(t *testing.T) {
	type fields struct {
		parent *node
	}
	type args struct {
		k Key
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		wantSibling *node
		wantExists  bool
	}{
		{
			name: "no siblings",
			fields: fields{
				parent: &node{
					entries:  []*Entry{},
					children: []*node{},
				},
			},
			args:        args{k: 3},
			wantSibling: nil,
			wantExists:  false,
		},
		{
			name: "left sibling exists",
			fields: fields{
				parent: &node{
					entries: []*Entry{{3, nil}},
					children: []*node{
						{entries: []*Entry{{0, 0}, {1, 1}}},
						{entries: []*Entry{{3, 3}, {4, 4}}},
					},
				},
			},
			args:        args{k: 4},
			wantSibling: &node{entries: []*Entry{{0, 0}, {1, 1}}},
			wantExists:  true,
		},
		{
			name: "left sibling with parent on edge",
			fields: fields{
				parent: &node{
					entries: []*Entry{{3, nil}},
					children: []*node{
						{entries: []*Entry{{0, 0}, {1, 1}}},
						{entries: []*Entry{{3, 3}, {4, 4}}},
					},
				},
			},
			args:        args{k: 3},
			wantSibling: &node{entries: []*Entry{{0, 0}, {1, 1}}},
			wantExists:  true,
		},
		{
			name: "no left siblings, node is already leftmost",
			fields: fields{
				parent: &node{
					entries: []*Entry{{3, nil}},
					children: []*node{
						{entries: []*Entry{{0, 0}, {1, 1}}},
						{entries: []*Entry{{3, 3}, {4, 4}}},
					},
				},
			},
			args:        args{k: 0},
			wantSibling: nil,
			wantExists:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				parent: tt.fields.parent,
			}

			gotSibling, gotExists := n.leftSibling(tt.args.k)
			assert.Equal(t, tt.wantExists, gotExists)
			assert.Equal(t, tt.wantSibling, gotSibling)
		})
	}
}

func Test_node_rightSibling(t *testing.T) {
	type fields struct {
		parent *node
	}
	type args struct {
		k Key
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		wantSibling *node
		wantExists  bool
	}{
		{
			name: "no siblings",
			fields: fields{
				parent: &node{
					entries:  []*Entry{},
					children: []*node{},
				},
			},
			args:        args{k: 3},
			wantSibling: nil,
			wantExists:  false,
		},
		{
			name: "right sibling exists",
			fields: fields{
				parent: &node{
					entries: []*Entry{{3, nil}},
					children: []*node{
						{entries: []*Entry{{0, 0}, {1, 1}}},
						{entries: []*Entry{{3, 3}, {4, 4}}},
					},
				},
			},
			args:        args{k: 1},
			wantSibling: &node{entries: []*Entry{{3, 3}, {4, 4}}},
			wantExists:  true,
		},
		{
			name: "right sibling with parent on edge",
			fields: fields{
				parent: &node{
					entries: []*Entry{{3, nil}},
					children: []*node{
						{entries: []*Entry{{0, 0}, {1, 1}}},
						{entries: []*Entry{{3, 3}, {4, 4}}},
					},
				},
			},
			args:        args{k: 2},
			wantSibling: &node{entries: []*Entry{{3, 3}, {4, 4}}},
			wantExists:  true,
		},
		{
			name: "no right siblings, node is already rightmost",
			fields: fields{
				parent: &node{
					entries: []*Entry{{3, nil}},
					children: []*node{
						{entries: []*Entry{{0, 0}, {1, 1}}},
						{entries: []*Entry{{3, 3}, {4, 4}}},
					},
				},
			},
			args:        args{k: 3},
			wantSibling: nil,
			wantExists:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				parent: tt.fields.parent,
			}

			gotSibling, gotExists := n.rightSibling(tt.args.k)
			assert.Equal(t, tt.wantExists, gotExists)
			assert.Equal(t, tt.wantSibling, gotSibling)
		})
	}
}

// Helper functions

// Repairs all the parent pointers throughout a tree
func repairParents(t *testing.T, tree *node, parent *node) *node {
	t.Helper()

	tree.parent = parent
	for i := range tree.children {
		repairParents(t, tree.children[i], tree)
	}

	return tree
}

func TestBtree_Height(t *testing.T) {
	type fields struct {
		root *node
	}

	root := repairParents(t, &node{
		entries: nil,
		children: []*node{
			{
				children: []*node{
					{
						children: []*node{
							{
								children: []*node{
									{
										entries: []*Entry{{1, 1}},
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil)

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"full tree", fields{root}, 4},
		{"partial tree", fields{root.children[0]}, 3},
		{"almost at bottom of tree", fields{root.children[0].children[0].children[0]}, 1},
		{"bottom of tree", fields{root.children[0].children[0].children[0].children[0]}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Btree{
				root: tt.fields.root,
			}

			assert.Equal(t, tt.want, b.Height())
		})
	}
}
