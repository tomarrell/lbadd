package schema

import "github.com/tomarrell/lbadd/internal/database/table"

// Schema describes a schema, which consists of zero or more tables.
type Schema interface {
	Table(name string) (table.Table, bool)
}
