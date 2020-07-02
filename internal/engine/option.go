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

// WithTimeProvider sets a time provider, which will be used by the engine to
// evaluate expressions, that require a timestamp, such as the function NOW().
func WithTimeProvider(tp timeProvider) Option {
	return func(e *Engine) {
		e.timeProvider = tp
	}
}

// WithRandomProvider sets a random provider, which will be used by the engine
// to evaluate expressions, that require a random source, such as the function
// RANDOM().
func WithRandomProvider(rp randomProvider) Option {
	return func(e *Engine) {
		e.randomProvider = rp
	}
}
