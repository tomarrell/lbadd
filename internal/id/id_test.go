package id_test

import (
	"testing"

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
