package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_node_depth(t *testing.T) {
	type fields struct {
		parent *node
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
		{
			name:   "root",
			fields: fields{parent: root},
			want:   0,
		},
		{
			name:   "one level down",
			fields: fields{parent: root.children[0]},
			want:   1,
		},
		{
			name:   "two levels down",
			fields: fields{parent: root.children[0].children[0]},
			want:   2,
		},
		{
			name:   "three levels down",
			fields: fields{parent: root.children[0].children[0].children[0]},
			want:   3,
		},
		{
			name:   "four levels down",
			fields: fields{parent: root.children[0].children[0].children[0].children[0]},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				parent: tt.fields.parent,
			}

			assert.Equal(t, tt.want, n.depth())
		})
	}
}
