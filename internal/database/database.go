package database

import "github.com/tomarrell/lbadd/internal/database/storage"

import "github.com/tomarrell/lbadd/internal/btree"

type key = btree.Key
type value = btree.Value

// storage defines the interface to be implemented by
// the b-tree
type storage interface {
	Get(k key) (v *btree.Entry, exists bool)
	Insert(k key, v value)
	Remove(k key) (removed bool)
	GetAll(limit int) []*btree.Entry
	GetAbove(k key, limit int) []*btree.Entry
	GetBelow(k key, limit int) []*btree.Entry
	GetBetween(low, high key, limit int) []*btree.Entry
}

type table struct {
	name    string
	store   storage.Storage
	columns []column
}

type db struct {
	tables map[string]table
}

func newDB() *db {
	tables := make(map[string]table)

	return &db{
		tables: tables,
	}
}
