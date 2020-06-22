package storage

import "github.com/rs/zerolog"

type Option func(*DBFile)

func WithCacheSize(size int) Option {
	return func(db *DBFile) {
		db.cacheSize = size
	}
}

func WithLogger(log zerolog.Logger) Option {
	return func(db *DBFile) {
		db.log = log
	}
}
