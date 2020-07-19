package storage

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

// Config is an intermediate layer to interact with configuration keys and
// values of a database file. It holds a pointer to the database file, so this
// struct may be copied and re-used.
type Config struct {
	db *DBFile
}

// Config returns an intermediate layer to interact with the keys and values
// stored in the config page of the database file. The returned struct may be
// copied and re-used.
func (db *DBFile) Config() Config {
	return Config{db}
}

// GetString returns the value associated with the given key, or an error, if
// there is no such value.
func (c Config) GetString(key string) (string, error) {
	cell, ok := c.db.configPage.Cell([]byte(key))
	if !ok {
		return "", ErrNoSuchConfigKey
	}
	if cell.Type() != page.CellTypeRecord {
		return "", fmt.Errorf("expected cell '%v' to be a record cell, but was %v", key, cell.Type())
	}
	return string(cell.(page.RecordCell).Record), nil
}

// SetString associates the given value with the given key. If there already is
// such a key present in the config, the value will be overwritten.
func (c Config) SetString(key, value string) error {
	err := c.db.configPage.StoreRecordCell(page.RecordCell{
		Key:    []byte(key),
		Record: []byte(value),
	})
	if err != nil {
		return fmt.Errorf("store record cell: %w", err)
	}
	c.db.configPage.MarkDirty()
	return nil
}
