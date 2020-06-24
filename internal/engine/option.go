package engine

import (
	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/engine/profile"
)

// Option is an option that can is applied to an Engine on creation.
type Option func(*Engine)

// WithLogger specifies a logger for the Engine.
func WithLogger(log zerolog.Logger) Option {
	return func(e *Engine) {
		e.log = log
	}
}

// WithProfiler passes a profiler into the engine. The default for the engine is
// not using a profiler at all.
func WithProfiler(profiler *profile.Profiler) Option {
	return func(e *Engine) {
		e.profiler = profiler
	}
}

// WithBuiltinFunction registeres or overwrites an existing builtin function.
// Use this to (e.g.) overwrite the SQL function NOW() to get constant
// timestamps.
func WithBuiltinFunction(name string, fn builtinFunction) Option {
	return func(e *Engine) {
		e.builtinFunctions[name] = fn
	}
}
