package lbadd

type table struct {
	name    string
	store   *storage
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
