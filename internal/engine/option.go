package engine

import "github.com/rs/zerolog"

// Option is an option that can is applied to an Engine on creation.
type Option func(*Engine)

// WithLogger specifies a logger for the Engine.
func WithLogger(log zerolog.Logger) Option {
	return func(e *Engine) {
		e.log = log
	}
}
