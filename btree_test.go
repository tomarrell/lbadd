package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTree(t *testing.T) {
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
				assert.Equal(t, g.value, bt.get(g.key))
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
