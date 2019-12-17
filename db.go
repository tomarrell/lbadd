package main

type table struct {
	store *storage
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
