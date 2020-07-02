package engine

import (
	"math/rand"
	"time"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/profile"
	"github.com/tomarrell/lbadd/internal/engine/result"
	"github.com/tomarrell/lbadd/internal/engine/storage"
	"github.com/tomarrell/lbadd/internal/engine/storage/cache"
)

type timeProvider func() time.Time
type randomProvider func() int64

// Engine is the component that is used to evaluate commands.
type Engine struct {
	log       zerolog.Logger
	dbFile    *storage.DBFile
	pageCache cache.Cache
	profiler  *profile.Profiler

	timeProvider   timeProvider
	randomProvider randomProvider
}

// New creates a new engine object and applies the given options to it.
func New(dbFile *storage.DBFile, opts ...Option) (Engine, error) {
	e := Engine{
		log:       zerolog.Nop(),
		dbFile:    dbFile,
		pageCache: dbFile.Cache(),

		timeProvider:   time.Now,
		randomProvider: func() int64 { return int64(rand.Uint64()) },
	}
	for _, opt := range opts {
		opt(&e)
	}
	return e, nil
}

// Evaluate evaluates the given command. This may mutate the state of the
// database, and changes may occur to the database file.
func (e Engine) Evaluate(cmd command.Command) (result.Result, error) {
	defer e.profiler.Enter(EvtEvaluate).Exit()

	_ = e.eq
	_ = e.lt
	_ = e.gt
	_ = e.lteq
	_ = e.gteq

	ctx := ExecutionContext{
		executionContext: &executionContext{},
	}
	return e.evaluate(ctx, cmd)
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
