package scanner

import (
	"testing"
)

func Test_hasNext(t *testing.T) {
	input := "SELECT "
	scanner := New([]rune(input))
	for scanner.HasNext() {
		scanner.Next()
	}
}
