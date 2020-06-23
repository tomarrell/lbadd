package engine

import (
	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/storage"
	"github.com/tomarrell/lbadd/internal/engine/storage/cache"
)

// Engine is the component that is used to evaluate commands.
type Engine struct {
	log       zerolog.Logger
	dbFile    *storage.DBFile
	pageCache cache.Cache
}

// New creates a new engine object and applies the given options to it.
func New(dbFile *storage.DBFile, opts ...Option) (*Engine, error) {
	e := &Engine{
		log:       zerolog.Nop(),
		dbFile:    dbFile,
		pageCache: dbFile.Cache(),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e, nil
}

// Evaluate evaluates the given command. This may mutate the state of the
// database, and changes may occur to the database file.
func (e Engine) Evaluate(cmd command.Command) (Result, error) {
	_ = e.eq
	_ = e.lt
	_ = e.gt
	_ = e.lteq
	_ = e.gteq
	return nil, nil
}

// Closed determines whether the underlying database file was closed. If so,
// this engine is considered closed, as it can no longer operate on the
// underlying file.
func (e Engine) Closed() bool {
	return e.dbFile.Closed()
}

// Close closes the underlying database file.
func (e Engine) Close() error {
	return e.dbFile.Close()
}
