package table

import (
	"github.com/tomarrell/lbadd/internal/database/column"
	"github.com/tomarrell/lbadd/internal/database/storage"
)

// Table describes a table that consists of a schema, a name, columns and a
// storage component, that is used to store the table's data.
type Table interface {
	Schema() string
	Name() string
	Columns() []column.Column
	Storage() storage.Storage
}
