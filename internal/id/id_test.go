package id_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/id"
)

func TestIDThreadSafe(t *testing.T) {
	// This is sufficient for the race detector to detect a race if Create() is
	// not safe for concurrent use.
	for i := 0; i < 5; i++ {
		go func() {
			_ = id.Create()
		}()
	}
}

func TestIDEquality(t *testing.T) {
	assert := assert.New(t)

	id1 := id.Create()
	id2, err := id.Parse(id1.Bytes())
	assert.NoError(err)

	assert.Equal(id1, id2)
	assert.True(id1 == id2)
}
