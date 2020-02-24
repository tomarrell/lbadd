package database

import "github.com/tomarrell/lbadd/internal/database/table"

// DB describes a database.
type DB interface {
	Table(schema, name string) (table.Table, bool)
}
