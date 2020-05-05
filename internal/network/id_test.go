package network

import "testing"

func TestIDThreadSafe(t *testing.T) {
	// This is sufficient for the race detector to detect a race if createID is
	// not safe for concurrent use.
	for i := 0; i < 5; i++ {
		go func() {
			_ = createID()
		}()
	}
}
