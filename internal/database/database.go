package database

import "github.com/tomarrell/lbadd/internal/database/storage"

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
