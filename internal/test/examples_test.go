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

func TestExample03(t *testing.T) {
	RunAndCompare(t, Test{
		Name:      "example03",
		Statement: `SELECT * FROM (VALUES (1, 2, 3), (4, 5, 6), (7, 5, 9))`,
	})
}

func TestExample04(t *testing.T) {
	RunAndCompare(t, Test{
		Name:      "example04",
		Statement: `SELECT * FROM (VALUES (1, 2, 3), (4, 5, 6), (7, 5, 9)) WHERE column2 = 5`,
	})
}
