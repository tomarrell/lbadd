package storage

import "github.com/rs/zerolog"

// Option is an option that can is applied to a DBFile on creation.
type Option func(*DBFile)

// WithCacheSize specifies a cache size for the DBFile.
func WithCacheSize(size int) Option {
	return func(db *DBFile) {
		db.cacheSize = size
	}
}

// WithLogger specifies a logger for the DBFile.
func WithLogger(log zerolog.Logger) Option {
	return func(db *DBFile) {
		db.log = log
	}
}
