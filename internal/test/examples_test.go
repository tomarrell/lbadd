package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/engine"
)

func TestExample01(t *testing.T) {
	RunAndCompare(t, Test{
		Name:      "example01",
		Statement: `VALUES (RANDOM())`,
		EngineOptions: []engine.Option{
			engine.WithRandomProvider(func() int64 { return 85734726843 }),
		},
	})
}

func TestExample02(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2020-07-02T14:03:27Z")
	assert.NoError(t, err)

	RunAndCompare(t, Test{
		Name:      "example02",
		Statement: `VALUES (NOW(), RANDOM())`,
		EngineOptions: []engine.Option{
			engine.WithTimeProvider(func() time.Time { return timestamp }),
			engine.WithRandomProvider(func() int64 { return 85734726843 }),
		},
	})
}
