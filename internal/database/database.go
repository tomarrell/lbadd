package database

import "github.com/tomarrell/lbadd/internal/database/schema"

// DB describes a database.
type DB interface {
	Schema(name string) (schema.Schema, bool)
}
