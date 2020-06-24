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

func WithProfiler(profiler *profile.Profiler) Option {
	return func(e *Engine) {
		e.profiler = profiler
	}
}
