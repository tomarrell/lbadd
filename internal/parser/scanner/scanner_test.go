package scanner

import (
	"testing"
)

func Test_hasNext(t *testing.T) {
	t.SkipNow()

	input := "SELECT "
	scanner := New([]rune(input))
	for scanner.HasNext() {
		scanner.Next()
	}
}
